package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"strings"
)

const (
	resumeIndex = "resumes"

	indexTag                 = "index"
	filterParamsTag          = "filter_params"
	sortParamsTag            = "sort_params"
	singleStringFilterParams = "*FilterParams_SingleString"
	singleIntFilterParams    = "*FilterParams_SingleInt"
	arrayIntFilterParams     = "*FilterParams_IntArray"
	sortParamTypeASC         = "*SortParams_ASC"
	sortParamTypeDESC        = "*SortParams_DESC"
)

//go:generate go run gen.go
func main() {
	goPath := os.Getenv("GOPATH")
	pkgPath := goPath + "/src/github.com/Stafford1986/test_manticore/pb"
	pkg := loadPackage(pkgPath)

	objResume := pkg.Types.Scope().Lookup("ResumeEntity")
	if objResume == nil {
		log.Fatal("not found in declared types of")
	}

	if _, ok := objResume.(*types.TypeName); !ok {
		log.Fatalf("%v is not a named type", objResume)
	}

	structTypeResume, ok := objResume.Type().Underlying().(*types.Struct)
	if !ok {
		log.Fatalf("type %v is not a struct", objResume)
	}

	err := generateResumeSaveMethods("pb", pkgPath, "ResumeEntity", structTypeResume)
	if err != nil {
		log.Fatal(err)
	}

	objResumeSearch := pkg.Types.Scope().Lookup("ResumeSearchEntity")
	if objResumeSearch == nil {
		log.Fatal("not found in declared types of resume search entity")
	}

	if _, ok := objResumeSearch.(*types.TypeName); !ok {
		log.Fatalf("%v is not a named type", objResumeSearch)
	}

	structTypeResumeSearch, ok := objResumeSearch.Type().Underlying().(*types.Struct)
	if !ok {
		log.Fatalf("type %v is not a struct", objResumeSearch)
	}

	err = generateResumeFindMethods("pb", pkgPath, "ResumeSearchEntity", structTypeResume, structTypeResumeSearch)
	if err != nil {
		log.Fatal(err)
	}

	err = generateResumeConvertMethod("pb", pkgPath, "ResumeEntity", structTypeResume)
	if err != nil {
		log.Fatal(err)
	}
	err = generateResumeVet("pb", pkgPath, "ResumeEntity", structTypeResume)
	if err != nil {
		log.Fatal(err)
	}
}

func generateResumeSaveMethods(pkgName string, pkgPath string, entityName string, structType *types.Struct) error {
	f := jen.NewFile(pkgName)
	// Add a package comment, so IDEs detect files as generated
	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	methodInsert := "BuildInsertQuery"
	methodUpsert := "BuildUpsertQuery"

	var (
		headBlock   []jen.Code
		shareBlock  []jen.Code
		insertBlosc []jen.Code
		upsertBlosc []jen.Code
	)

	structFields := make([]string, 0, structType.NumFields())
	structFieldTags := make([]string, 0, structType.NumFields())
	structFieldType := make([]string, 0, structType.NumFields())

	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		tagValue := structType.Tag(i)
		if tagValue == "" {
			continue
		}

		tags := strings.Split(tagValue, " ")

		var (
			tagName string
		)
		for _, v := range tags {
			if strings.HasPrefix(v, "db") {
				st := strings.Split(v, ":")
				if len(st) != 2 {
					return fmt.Errorf("wrong field tag: %s", field.String())
				}
				tagName = st[1]
			}
		}
		if tagName == "" {
			return fmt.Errorf("wrong field tag: %s", field.String())
		}

		tagParts := strings.Split(tagName, ",")
		tagName = strings.ReplaceAll(tagParts[0], "\"", "")

		structFields = append(structFields, field.Name())
		structFieldType = append(structFieldType, field.Type().String())
		structFieldTags = append(structFieldTags, tagName)
	}

	headBlock = append(headBlock, jen.Var().Defs(
		jen.Id("resErr").Error(),
		jen.Id("resValues").Op("=").Make(jen.Index().String(), jen.Lit(0), jen.Lit(len(structFields)))),
	)

	insertBlosc = append(insertBlosc, jen.Id("sb").
		Op(":=").
		Qual("bytes", "NewBufferString").
		Call(jen.Lit("INSERT INTO "+resumeIndex+"("+strings.Join(structFieldTags, ", ")+")"+" VALUES(")))

	upsertBlosc = append(upsertBlosc, jen.Id("sb").
		Op(":=").
		Qual("bytes", "NewBufferString").
		Call(jen.Lit("REPLACE INTO "+resumeIndex+"("+strings.Join(structFieldTags, ", ")+")"+" VALUES(")))

	for i, v := range structFields {
		tf := structFieldType[i]
		switch tf {
		case "int64", "uint32":
			shareBlock = append(shareBlock, jen.Id("resValues").
				Op("=").
				Append(jen.Id("resValues"), jen.Qual("strconv", "FormatInt").
					Call(jen.Id("int64").Parens(jen.Id("re").Dot(v)), jen.Id("10"))))

		case "[]uint32":
			shareBlock = append(shareBlock, jen.Id("resValues").
				Op("=").
				Append(jen.Id("resValues"), jen.Func().
					Parens(jen.Empty()).String().Block(
					jen.Id("sb").
						Op(":=").
						Qual("bytes", "NewBufferString").
						Parens(jen.Lit("(")),
					jen.If(jen.Len(jen.Id("re").Dot(v)).Op("==").Id("0").Block(
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Lit("0)")),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
						jen.Return(jen.Id("sb").Dot("String()")),
					)),
					jen.For(jen.Id("i").Id(",").Id("v").Op(":=").Range().Id("re").Dot(v).Block(
						jen.If(jen.Id("i").Op("==").Len(jen.Id("re").Dot(v)).Op("-").Id("1")).Block(
							jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
								Parens(jen.Qual("strconv", "FormatInt").
									Call(jen.Int64().Parens(jen.Id("v")), jen.Id("10"))),
							jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
								Parens(jen.Lit(")")),
							jen.If(jen.Id("err").Op("!= nil")).Block(
								jen.Id("resErr").Op("=").Qual(
									"github.com/hashicorp/go-multierror", "Append").
									Call(jen.Id("resErr"), jen.Id("err")),
							),
							jen.Break(),
						),
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Qual("strconv", "FormatInt").
								Call(jen.Int64().Parens(jen.Id("v")), jen.Id("10"))),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(",")),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
					),
						jen.Return(jen.Id("sb").Dot("String()")),
					),
				).Call()))
		case "bool":
			shareBlock = append(shareBlock, jen.Id("resValues").
				Op("=").
				Append(jen.Id("resValues"), jen.Func().Parens(jen.Empty()).String().Block(
					jen.If(jen.Id("re").Dot(v).Op("==").True().Block(
						jen.Return(jen.Lit("1"))),
						jen.Return(jen.Lit("0")),
					),
				).
					Call()))
		case "string":
			shareBlock = append(shareBlock, jen.Id("resValues").
				Op("=").
				Append(jen.Id("resValues"), jen.Func().
					Parens(jen.Empty()).String().Block(
					jen.Id("sb").
						Op(":=").
						Qual("bytes", "NewBufferString").
						Parens(jen.Lit("'")),
					jen.If(jen.Len(jen.Id("re").Dot(v)).Op("==").Id("0").Block(
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" '")),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
						jen.Return(jen.Id("sb").Dot("String()")),
					)),

					jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
						Parens(jen.Id("re").Dot(v)),
					jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
						Parens(jen.Lit("'")),

					jen.If(jen.Id("err").Op("!= nil")).Block(
						jen.Id("resErr").Op("=").Qual(
							"github.com/hashicorp/go-multierror", "Append").
							Call(jen.Id("resErr"), jen.Id("err")),
					),

					jen.Return(jen.Id("sb").Dot("String()")),
				).Call()))
		default:
			jen.Id("resErr").Op("=").Qual(
				"github.com/hashicorp/go-multierror", "Append").
				Call(jen.Id("resErr"), jen.Qual(
					"errors", "New").Call(jen.Lit("err. got wrong type")))
		}
	}

	shareBlock = append(shareBlock,
		jen.Id("_, err").Op(":=").Id("sb").
			Dot("WriteString").
			Parens(jen.Qual("strings", "Join").
				Call(jen.Id("resValues"), jen.Lit(", "))))
	shareBlock = append(shareBlock,
		jen.Id("_, err").Op("=").Id("sb").
			Dot("WriteString").Call(jen.Lit(");")))

	shareBlock = append(shareBlock, jen.If(jen.Id("err").Op("!= nil")).Block(
		jen.Id("resErr").Op("=").Qual(
			"github.com/hashicorp/go-multierror", "Append").
			Call(jen.Id("resErr"), jen.Id("err")),
	))

	shareBlock = append(shareBlock, jen.If(jen.Id("resErr").Op("!= nil")).Block(
		jen.Return(jen.Lit(""), jen.Id("resErr")),
	))

	shareBlock = append(shareBlock, jen.Return(jen.Id("sb").Dot("String()"), jen.Nil()))

	finalInsertBlock := make([]jen.Code, 0, len(headBlock)+len(insertBlosc)+len(shareBlock))
	finalInsertBlock = append(finalInsertBlock, headBlock...)
	finalInsertBlock = append(finalInsertBlock, insertBlosc...)
	finalInsertBlock = append(finalInsertBlock, shareBlock...)
	f.Func().Params(jen.Id("re").Op("*").Id(entityName)).Id(methodInsert).Params().Params(jen.String(), jen.Error()).Block(finalInsertBlock...)

	finalUpsertBlock := make([]jen.Code, 0, len(headBlock)+len(upsertBlosc)+len(shareBlock))
	finalUpsertBlock = append(finalUpsertBlock, headBlock...)
	finalUpsertBlock = append(finalUpsertBlock, upsertBlosc...)
	finalUpsertBlock = append(finalUpsertBlock, shareBlock...)
	f.Func().Params(jen.Id("re").Op("*").Id(entityName)).Id(methodUpsert).Params().Params(jen.String(), jen.Error()).Block(finalUpsertBlock...)

	file_name := pkgPath + "/create_resume_helper_gen.go"

	return f.Save(file_name)
}

func generateResumeFindMethods(pkgName string, pkgPath string, entityName string, targetStruct *types.Struct, searchStructType *types.Struct) error {
	f := jen.NewFile(pkgName)
	// Add a package comment, so IDEs detect files as generated
	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	methodBuildSearch := "BuildSearchQuery"

	var (
		headBlock  []jen.Code
		shareBlock []jen.Code
	)

	targetStructFields := make([]string, 0, targetStruct.NumFields())
	targetStructFieldTags := make([]string, 0, targetStruct.NumFields())
	targetFieldTagMap := make(map[string]string, targetStruct.NumFields())
	targetStructFieldType := make([]string, 0, targetStruct.NumFields())

	for i := 0; i < targetStruct.NumFields(); i++ {
		field := targetStruct.Field(i)
		tagValue := targetStruct.Tag(i)
		if tagValue == "" {
			continue
		}

		tags := strings.Split(tagValue, " ")

		var (
			tagName string
			tagType string
		)
		for _, v := range tags {
			if strings.HasPrefix(v, "db") {
				st := strings.Split(v, ":")
				if len(st) != 2 {
					return fmt.Errorf("wrong field tag: %s", field.String())
				}
				tagName = st[1]
			}
		}
		if tagName == "" {
			return fmt.Errorf("wrong field tag: %s", field.String())
		}

		tagParts := strings.Split(tagName, ",")
		tagName = strings.ReplaceAll(tagParts[0], "\"", "")
		if len(tagParts) < 2 {
			tagType = ""
		} else {
			tagType = strings.ReplaceAll(tagParts[1], "\"", "")
		}

		targetStructFields = append(targetStructFields, field.Name())
		targetFieldTagMap[tagName] = tagType
		targetStructFieldTags = append(targetStructFieldTags, tagName)
		targetStructFieldType = append(targetStructFieldType, field.Type().String())
	}

	searchStructFields := make([]string, 0, searchStructType.NumFields())
	searchStructTags := make([]string, 0, searchStructType.NumFields())
	searchStructFieldType := make([]string, 0, searchStructType.NumFields())
	queryFieldIndex := 0

	for i := 0; i < searchStructType.NumFields(); i++ {
		field := searchStructType.Field(i)
		tagValue := searchStructType.Tag(i)
		if tagValue == "" {
			continue
		}

		tags := strings.Split(tagValue, " ")

		var (
			tagName string
		)
		for i, v := range tags {
			if strings.HasPrefix(v, "json") {
				st := strings.Split(v, ":")
				if len(st) != 2 {
					return fmt.Errorf("wrong field tag search entiry: %s", field.String())
				}
				tagName = st[1]
				if tagName == "query" {
					queryFieldIndex = i
				}
			}
		}
		if tagName == "" {
			return fmt.Errorf("wrong field tag: %s", field.String())
		}

		tagParts := strings.Split(tagName, ",")
		tagName = strings.ReplaceAll(tagParts[0], "\"", "")

		searchStructFields = append(searchStructFields, field.Name())
		searchStructTags = append(searchStructTags, tagName)
		searchStructFieldType = append(searchStructFieldType, field.Type().String())
	}

	f.Line().Var().Defs(
		jen.Id("targetResumeFieldMap").Op("=").Make(jen.Map(jen.String()).String(), jen.Lit(len(targetFieldTagMap))),
		jen.Id("targetResumeFieldMapType").Op("=").Make(jen.Map(jen.String()).String(), jen.Lit(len(targetFieldTagMap))),
	)

	shareBlock = append(shareBlock, jen.Var().Defs(
		jen.Id("resErr").Error(),
	))

	for k, v := range targetFieldTagMap {
		headBlock = append(headBlock,
			jen.Id("targetResumeFieldMap").Index(jen.Lit(k)).Op("=").Lit(v))
	}

	for i, v := range targetStructFieldType {
		headBlock = append(headBlock,
			jen.Id("targetResumeFieldMapType").Index(jen.Lit(targetStructFieldTags[i])).Op("=").Lit(v))
	}

	f.Line().Func().Id("init").Call().Block(headBlock...)

	shareBlock = append(shareBlock, jen.If(jen.Len(jen.Id("re").Dot(searchStructFields[queryFieldIndex])).Op("==").Id("0").Block(
		jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. empty search req"))),
	)))

	shareBlock = append(shareBlock, jen.Id("query").Op(":=").Id("re").Dot(searchStructFields[queryFieldIndex]))

	shareBlock = append(shareBlock, jen.Id("sb").
		Op(":=").
		Qual("bytes", "NewBufferString").
		Call(jen.Qual("fmt", "Sprintf").Call(jen.Lit("SELECT * FROM "+resumeIndex+" WHERE MATCH('*%s*')"), jen.Id("query"))))

	for i, v := range searchStructFields {
		if i == queryFieldIndex {
			continue
		}
		fildTag := searchStructTags[i]
		switch fildTag {
		case filterParamsTag:
			shareBlock = append(shareBlock, jen.For(jen.Id("_").Id(",").Id("v").Op(":=").Range().Id("re").Dot(v).Block(
				jen.If(jen.Id("v").Dot("Filter")).Op("==").Lit("").Block(
					jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. empty filter field"))),
				),
				jen.Id("fv").Id(",").Id("ok").Op(":=").Id("targetResumeFieldMap").Index(jen.Id("v").Dot("Filter")),
				jen.If(jen.Id("!ok")).Block(
					jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. wrong filter param name"))),
				),
				jen.If(jen.Id("fv")).Op("!=").Lit("filter").Block(
					jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. field can't be used as filter"))),
				),
				jen.Switch(jen.Id("v").Dot("Value").Dot("(type)").Block(
					jen.Case(jen.Id(singleIntFilterParams)).Block(
						jen.Id("ft").Id(",").Id("ok").Op(":=").Id("targetResumeFieldMapType").Index(jen.Id("v").Dot("Filter")),
						jen.If(jen.Id("!ok")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. can't find filter name"))),
						),
						jen.If(jen.Id("ft").Op("!=").Lit("uint32")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. mismatch filter value type"))),
						),
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" AND ")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Id("v").Dot("Filter")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" = ")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Qual("strconv", "FormatInt").Call(jen.Int64().Parens(jen.Id("v").Dot("GetSingleInt()")), jen.Id("10"))),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
					),
					jen.Case(jen.Id(singleStringFilterParams)).Block(
						jen.Id("ft").Id(",").Id("ok").Op(":=").Id("targetResumeFieldMapType").Index(jen.Id("v").Dot("Filter")),
						jen.If(jen.Id("!ok")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. can't find filter name"))),
						),
						jen.If(jen.Id("ft").Op("!=").Lit("string")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. mismatch filter value type"))),
						),
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" AND ")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Id("v").Dot("Filter")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" = ")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Id("v").Dot("GetSingleString()")),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
					),
					jen.Case(jen.Id(arrayIntFilterParams)).Block(
						jen.Id("ft").Id(",").Id("ok").Op(":=").Id("targetResumeFieldMapType").Index(jen.Id("v").Dot("Filter")),
						jen.If(jen.Id("!ok")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. can't find filter name"))),
						),
						jen.If(jen.Id("ft").Op("!=").Lit("[]uint32")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. mismatch filter value type"))),
						),
						jen.Id("arr").Op(":=").Id("v").Dot("GetIntArray()"),
						jen.If(jen.Id("arr").Op("==").Id("nil")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. got nil array"))),
						),

						jen.Id("arrVal").Op(":=").Id("arr").Dot("GetValues()"),
						jen.If(jen.Len(jen.Id("arrVal")).Op("==").Id("0")).Block(
							jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. got empty array"))),
						),

						jen.Var().Defs(
							jen.Id("sbp").Op("=").Qual("bytes", "NewBufferString").Call(jen.Lit("")),
						),

						jen.For(jen.Id("i, v").Op(":=").Range().Id("arrVal").Block(
							jen.Var().Id("err").Error(),
							jen.If(jen.Id("i").Op("==").Id("0")).Block(
								jen.Id("_, err").Op("=").Id("sbp").Dot("WriteString").Call(jen.Lit("(")),
							),
							jen.Id("_, err").Op("=").Id("sbp").Dot("WriteString").
								Parens(jen.Qual("strconv", "FormatInt").Call(jen.Int64().Parens(jen.Id("v")), jen.Id("10"))),
							jen.If(jen.Id("i").Op("==").Len(jen.Id("arrVal")).Id("- 1")).Block(
								jen.Id("_, err").Op("=").Id("sbp").Dot("WriteString").Call(jen.Lit(")")),
							).Else().Block(
								jen.Id("_, err").Op("=").Id("sbp").Dot("WriteString").Call(jen.Lit(",")),
							),
							jen.If(jen.Id("err").Op("!= nil")).Block(
								jen.Id("resErr").Op("=").Qual(
									"github.com/hashicorp/go-multierror", "Append").
									Call(jen.Id("resErr"), jen.Id("err")),
							),
						))),

					jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
						Parens(jen.Lit(" AND ")),
					jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
						Parens(jen.Id("v").Dot("Filter")),
					jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
						Parens(jen.Lit(" IN")),
					jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
						Parens(jen.Id("sbp").Dot("String()")),
					jen.If(jen.Id("err").Op("!= nil")).Block(
						jen.Id("resErr").Op("=").Qual(
							"github.com/hashicorp/go-multierror", "Append").
							Call(jen.Id("resErr"), jen.Id("err")),
					),
				),
				),
			),
			))
		case sortParamsTag:
			shareBlock = append(shareBlock, jen.For(jen.Id("_").Id(",").Id("v").Op(":=").Range().Id("re").Dot(v).Block(
				jen.If(jen.Id("v").Dot("Field")).Op("==").Lit("").Block(
					jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. empty sort field"))),
				),
				jen.Id("fv").Id(",").Id("ok").Op(":=").Id("targetResumeFieldMap").Index(jen.Id("v").Dot("Field")),
				jen.If(jen.Id("!ok")).Block(
					jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. wrong sort param name"))),
				),
				jen.If(jen.Id("fv")).Op("!=").Lit("").Block(
					jen.Return(jen.Lit(""), jen.Qual("errors", "New").Call(jen.Lit("err validation. field can't be used as sorted"))),
				),
				jen.Switch(jen.Id("v").Dot("Order").Dot("(type)").Block(
					jen.Case(jen.Id(sortParamTypeASC)).Block(
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" ORDER BY ")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Id("v").Dot("Field")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" ASC")),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
					),
					jen.Case(jen.Id(sortParamTypeDESC)).Block(
						jen.Id("_, err").Op(":=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" ORDER BY ")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Id("v").Dot("Field")),
						jen.Id("_, err").Op("=").Id("sb").Dot("WriteString").
							Parens(jen.Lit(" DESC")),
						jen.If(jen.Id("err").Op("!= nil")).Block(
							jen.Id("resErr").Op("=").Qual(
								"github.com/hashicorp/go-multierror", "Append").
								Call(jen.Id("resErr"), jen.Id("err")),
						),
					),
				),
				),
			),
			))
		default:
			shareBlock = append(shareBlock, jen.Return(jen.Lit(""),
				jen.Qual("errors", "New").Call(jen.Lit("err validation. unknown tag name"))))
		}
	}

	shareBlock = append(shareBlock,
		jen.Id("_, err").Op(":=").Id("sb").
			Dot("WriteString").Call(jen.Qual(
			"fmt", "Sprintf").Call(jen.Lit(" LIMIT %d;"), jen.Id("limit"))),
		jen.If(jen.Id("err").Op("!= nil")).Block(
			jen.Id("resErr").Op("=").Qual(
				"github.com/hashicorp/go-multierror", "Append").
				Call(jen.Id("resErr"), jen.Id("err")),
		))

	shareBlock = append(shareBlock, jen.If(jen.Id("resErr").Op("!= nil")).Block(
		jen.Return(jen.Lit(""), jen.Id("resErr")),
	))

	shareBlock = append(shareBlock, jen.Return(jen.Id("sb").Dot("String()"), jen.Nil()))

	f.Func().Params(jen.Id("re").Op("*").Id(entityName)).Id(methodBuildSearch).Params(jen.Id("limit").Int()).Params(jen.String(), jen.Error()).Block(shareBlock...)

	file_name := pkgPath + "/search_resume_helper_gen.go"

	return f.Save(file_name)
}

func generateResumeConvertMethod(packageName string, pkgPath string, entityName string, targetStruct *types.Struct) error {
	f := jen.NewFile(packageName)
	// Add a package comment, so IDEs detect files as generated
	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	methodParseResult := "ParseDbResult"
	//
	var (
		shareBlock []jen.Code
	)
	//
	targetStructFields := make([]string, 0, targetStruct.NumFields())
	targetStructFieldTags := make([]string, 0, targetStruct.NumFields())
	targetStructFieldType := make([]string, 0, targetStruct.NumFields())

	for i := 0; i < targetStruct.NumFields(); i++ {
		field := targetStruct.Field(i)
		tagValue := targetStruct.Tag(i)
		if tagValue == "" {
			continue
		}

		tags := strings.Split(tagValue, " ")

		var (
			tagName string
		)
		for _, v := range tags {
			if strings.HasPrefix(v, "db") {
				st := strings.Split(v, ":")
				if len(st) != 2 {
					return fmt.Errorf("wrong field tag: %s", field.String())
				}
				tagName = st[1]
			}
		}
		if tagName == "" {
			return fmt.Errorf("wrong field tag: %s", field.String())
		}

		tagParts := strings.Split(tagName, ",")
		tagName = strings.ReplaceAll(tagParts[0], "\"", "")

		targetStructFields = append(targetStructFields, field.Name())
		targetStructFieldTags = append(targetStructFieldTags, tagName)
		targetStructFieldType = append(targetStructFieldType, field.Type().String())
	}

	shareBlock = append(shareBlock, jen.Var().Defs(
		jen.Id("res").Op("=").Id("&"+entityName+"{}"),
	))

	shareBlock = append(shareBlock, jen.For(jen.Id("k, v").Op(":=").Range().Id("m").Id("{")))
	shareBlock = append(shareBlock, jen.Switch(jen.Id("k").Id("{")))

	for i, v := range targetStructFields {
		shareBlock = append(shareBlock, jen.Case(jen.Lit(targetStructFieldTags[i])))
		shareBlock = append(shareBlock, jen.Id("val, ok").Op(":=").Id("v").Dot("([]byte)"))
		shareBlock = append(shareBlock, jen.If(jen.Id("!ok")).Block(
			jen.Return(jen.Id("nil").Id(",").Qual(
				"errors", "New").
				Call(jen.Lit("err convert "+targetStructFieldTags[i])),
			),
		))
		t := targetStructFieldType[i]
		switch t {
		case "string":
			shareBlock = append(shareBlock, jen.Id("res").Dot(v).Op("=").String().Call(jen.Id("val")))
		case "uint32":
			shareBlock = append(shareBlock, jen.Id("p, err").Op(":=").Qual("strconv", "ParseUint").
				Call(jen.String().Call(jen.Id("val")), jen.Id("10"), jen.Id("32")))
			shareBlock = append(shareBlock, jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
				jen.Return(jen.Id("nil").Id(",").Qual(
					"errors", "New").
					Call(jen.Lit("err convert value to "+v)))),
			)
			shareBlock = append(shareBlock, jen.Id("res").Dot(v).Op("=").Uint32().Call(jen.Id("p")))
		case "[]uint32":
			shareBlock = append(shareBlock, jen.Var().Defs(
				jen.Id("r").Id("[]uint32"),
			))
			shareBlock = append(shareBlock, jen.Id("s").
				Op(":=").Qual("strings", "Split").
				Call(jen.String().Call(jen.Id("val")), jen.Lit(",")))
			shareBlock = append(shareBlock, jen.For(jen.Id("_, sv").Op(":=").Range().Id("s").Block(
				jen.Id("p, err").Op(":=").Qual("strconv", "ParseUint").
					Call(jen.String().Call(jen.Id("sv")), jen.Id("10"), jen.Id("32")),
				jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
					jen.Return(jen.Id("nil").Id(",").Qual(
						"errors", "New").
						Call(jen.Lit("err convert value to "+v)))),
				jen.Id("r").Op("=").Append(jen.Id("r"), jen.Uint32().Call(jen.Id("p")))),
			),
			)
			shareBlock = append(shareBlock, jen.Id("res").Dot(v).Op("=").Id("r"))
		case "int64":
			shareBlock = append(shareBlock, jen.Id("p, err").Op(":=").Qual("strconv", "ParseInt").
				Call(jen.String().Call(jen.Id("val")), jen.Id("10"), jen.Id("64")))
			shareBlock = append(shareBlock, jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
				jen.Return(jen.Id("nil").Id(",").Qual(
					"errors", "New").
					Call(jen.Lit("err convert value to "+v)))),
			)
			shareBlock = append(shareBlock, jen.Id("res").Dot(v).Op("=").Id("p"))
		case "bool":
			shareBlock = append(shareBlock, jen.Id("p, err").
				Op(":=").Qual("strconv", "Atoi").
				Call(jen.String().Call(jen.Id("val"))))
			shareBlock = append(shareBlock, jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
				jen.Return(jen.Id("nil").Id(",").Qual(
					"errors", "New").
					Call(jen.Lit("err convert value to "+v)))),
			)
			shareBlock = append(shareBlock, jen.If(jen.Id("p").Op("==").Id("0")).Block(
				jen.Id("res").Dot(v).Op("=").Id("false"),
			).Else().Block(
				jen.Id("res").Dot(v).Op("=").Id("true"),
			))
		default:
			//return fmt.Errorf("err. unsupported type: %s", t)
		}

	}
	shareBlock = append(shareBlock, jen.Default())
	shareBlock = append(shareBlock, jen.Return(jen.Id("nil").Id(",").Qual(
		"fmt", "Errorf").
		Call(jen.Lit("unknown field %s"), jen.Id("k")),
	))

	shareBlock = append(shareBlock, jen.Id("}"))
	shareBlock = append(shareBlock, jen.Id("}"))

	shareBlock = append(shareBlock, jen.Return(jen.Id("res"), jen.Nil()))

	f.Func().Params(jen.Id("re").Op("*").Id(entityName)).Id(methodParseResult).
		Params(jen.Id("m").Map(jen.String()).Interface()).Params(jen.Id("*"+entityName), jen.Error()).Block(shareBlock...)

	file_name := pkgPath + "/convert_resume_helper_gen.go"

	return f.Save(file_name)
}

func loadPackage(path string) *packages.Package {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		log.Fatal("loading packages for inspection: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		log.Fatalf("pakages err: %d", packages.PrintErrors(pkgs))
	}

	return pkgs[0]
}

func generateResumeVet(pkgName string, pkgPath string, entityName string, structType *types.Struct) error {
	f := jen.NewFile(pkgName)
	// Add a package comment, so IDEs detect files as generated
	f.PackageComment("Code generated by generator, DO NOT EDIT.")

	methodVet := "Vet"

	var (
		shareBlock []jen.Code
	)

	structFields := make([]string, 0, structType.NumFields())
	structFieldTagsType := make([]string, 0, structType.NumFields())

	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		tagValue := structType.Tag(i)
		if tagValue == "" {
			continue
		}

		tags := strings.Split(tagValue, " ")

		var (
			tagName string
		)
		for _, v := range tags {
			if strings.HasPrefix(v, "db") {
				st := strings.Split(v, ":")
				if len(st) != 2 {
					return fmt.Errorf("wrong field tag: %s", field.String())
				}
				tagName = st[1]
			}
		}
		if tagName == "" {
			return fmt.Errorf("wrong field tag: %s", field.String())
		}

		structFields = append(structFields, field.Name())

		tagParts := strings.Split(tagName, ",")
		if len(tagParts) < 2 {
			structFieldTagsType = append(structFieldTagsType, "")
			continue
		}
		tagValue = strings.ReplaceAll(tagParts[1], "\"", "")
		if tagValue == indexTag {
			structFieldTagsType = append(structFieldTagsType, tagValue)
			continue
		}

		structFieldTagsType = append(structFieldTagsType, "")
	}

	shareBlock = append(shareBlock, jen.Var().Defs(
		jen.Id("fv").String(),
		jen.Id("fvb").Index().Byte(),
		jen.Id("resErr").Error(),
		jen.Id("res").Op("=").Id("&"+entityName+"{}"),
		jen.Id("reBreakers").Op("=").Qual("regexp", "MustCompile").
			Call(jen.Id("`<.*>`")),
		jen.Id("reSymbols").Op("=").Qual("regexp", "MustCompile").
			Call(jen.Lit("['\"<>/()\\[\\];:*{}!?=\\-+_~`$^&#№%\\\\]")),
		jen.Id("reSpaces").Op("=").Qual("regexp", "MustCompile").
			Call(jen.Id("`\\s+`")),
	))

	for i, v := range structFields {
		tf := structFieldTagsType[i]
		switch tf {
		case indexTag:
			shareBlock = append(shareBlock, jen.Id("fv").
				Op("=").
				Id("re").Dot(v))

			shareBlock = append(shareBlock, jen.Id("fv").
				Op("=").Id("reBreakers").Dot("ReplaceAllString").
				Call(jen.Id("fv"), jen.Lit("")))

			shareBlock = append(shareBlock, jen.Id("fvb").
				Op("=").Id("reSymbols").Dot("ReplaceAll").
				Call(jen.Index().Byte().Call(jen.Id("fv")), jen.Index().Byte().Call(jen.Lit(""))))

			shareBlock = append(shareBlock, jen.Id("fvb").
				Op("=").Id("reSpaces").Dot("ReplaceAll").
				Call(jen.Index().Byte().Call(jen.Id("fvb")), jen.Index().Byte().Call(jen.Lit(" "))))
			shareBlock = append(shareBlock, jen.Id("fv").
				Op("=").Qual("strings", "TrimSpace").Call(jen.String().Call(jen.Id("fvb"))))
			shareBlock = append(shareBlock, jen.Id("res").Dot(v).
				Op("=").
				Id("fv"))

		case "":
			shareBlock = append(shareBlock, jen.Id("res").Dot(v).
				Op("=").
				Id("re").Dot(v))
		default:
			jen.Id("resErr").Op("=").Qual(
				"github.com/hashicorp/go-multierror", "Append").
				Call(jen.Id("resErr"), jen.Qual(
					"errors", "New").Call(jen.Lit("err. got unknown tag type")))
		}
	}

	shareBlock = append(shareBlock, jen.If(jen.Id("resErr").Op("!= nil")).Block(
		jen.Return(jen.Nil(), jen.Id("resErr")),
	))

	shareBlock = append(shareBlock,
		jen.Return(jen.Id("res"), jen.Nil()),
	)

	f.Func().Params(jen.Id("re").Op("*").Id(entityName)).Id(methodVet).Params().Params(jen.Id("*"+entityName), jen.Error()).Block(shareBlock...)

	file_name := pkgPath + "/vet_resume_helper_gen.go"

	return f.Save(file_name)
}
