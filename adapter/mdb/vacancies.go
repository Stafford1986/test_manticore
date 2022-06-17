package mdb

import (
	"context"
	"github.com/Stafford1986/test_manticore/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VacancyRep struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *VacancyRep {
	return &VacancyRep{
		db: db,
	}
}

func (rep *VacancyRep) GetVacancies(ctx context.Context, limit int, skip int) ([]*pb.VacancyEntity, error) {
	var items []*pb.VacancyEntity

	lookupSkillsStage := bson.D{
		{"$lookup", bson.D{
			{"from", "ref_skills"},
			{"localField", "skills"},
			{"foreignField", "_id"},
			{"as", "skills"},
		}},
	}

	lookupAddressStage := bson.D{{"$lookup", bson.D{
		{"from", "addresses"},
		{"localField", "address"},
		{"foreignField", "_id"},
		{"as", "address"},
	}},
	}

	addCityField := bson.D{
		{"$addFields", bson.D{
			{"city", bson.D{
				{"$arrayElemAt", bson.A{
					"$address.city_id", 0,
				}},
			}}}},
	}

	addCityLevelField := bson.D{
		{"$addFields", bson.D{
			{"city_level", bson.D{
				{"$arrayElemAt", bson.A{
					"$address.city_level", 0,
				}},
			}}}},
	}

	addSkills := bson.D{{"$addFields", bson.D{
		{"skills", bson.D{
			{"$reduce", bson.D{
				{"input", "$skills"},
				{"initialValue", ""},
				{"in", bson.D{
					{"$concat", bson.A{
						"$$value", bson.D{
							{"$cond", bson.D{
								{"if", bson.D{
									{"$eq", bson.A{
										"$$value",
										"",
									},
									},
								},
								},
								{"then", ""},
								{"else", " "},
							}},
						},
						"$$this.skill_title",
					}},
				}}},
			},
		}}}}}

	addBrandField := bson.D{
		{"$addFields", bson.D{
			{"brand", bson.D{
				{"$concat", bson.A{
					"Brand test ",
					bson.D{
						{"$convert", bson.D{
							{"input", "$_id"},
							{"to", "string"},
						}}},
				},
				}},
			}}},
	}

	addWebsiteUrlField := bson.D{
		{"$addFields", bson.D{
			{"website_url", bson.D{
				{"$concat", bson.A{
					"Website test ",
					bson.D{
						{"$convert", bson.D{
							{"input", "$_id"},
							{"to", "string"},
						}}},
				},
				}},
			}}},
	}
	addLogoField := bson.D{
		{"$addFields", bson.D{
			{"logo", bson.D{
				{"$concat", bson.A{
					"Logo test ",
					bson.D{
						{"$convert", bson.D{
							{"input", "$_id"},
							{"to", "string"},
						}}},
				},
				}},
			}}},
	}

	addIndustryGroupField := bson.D{
		{"$addFields", bson.D{
			{"industry_groups", bson.A{
				bson.D{
					{"$convert", bson.D{
						{"input", 1},
						{"to", "int"},
					}}},
				bson.D{{"$convert", bson.D{
					{"input", 2},
					{"to", "int"},
				}}},
				bson.D{{"$convert", bson.D{
					{"input", 3},
					{"to", "int"},
				}}},
			}}}},
	}

	skipStage := bson.D{
		{"$skip", skip},
	}

	limitStage := bson.D{
		{"$limit", limit},
	}

	project := bson.D{{"$project", bson.D{
		{"_id", 1},
		{"name", 1},
		{"active", 1},
		{"created_at", 1},
		{"updated_at", 1},
		{"skills", 1},
		{"city", 1},
		{"city_level", 1},
		{"brand", 1},
		{"website_url", 1},
		{"logo", 1},
		{"industry_groups", 1},
		{"company_id", 1},
		{"specialization", 1},
		{"metro", 1},
		{"salary_before_tax", 1},
		{"salary_curr", 1},
		{"job_responsibility", 1},
		{"job_requirement", 1},
		{"work_condition", 1},
		{"city_visibility", 1},
		{"vacancy_language", 1},
		{"business_trips", 1},
		{"self_employed", 1},
		{"ip_employed", 1},
		{"payment_period", 1},
		{"salary_from", 1},
		{"salary_to", 1},
		{"default_work_type", 1},
		{"work_type", 1},
		{"experience", 1},
		{"min_customer_languages", 1},
		{"driver_license", 1},
		{"driver_exp", 1},
		{"have_car", 1},
		{"restrictions", 1},
		{"list_respond_button", 1},
	}}}

	aggregateVal, err := rep.db.Aggregate(ctx, mongo.Pipeline{
		lookupSkillsStage,
		lookupAddressStage,
		addCityField,
		addCityLevelField,
		addSkills,
		addBrandField,
		addWebsiteUrlField,
		addLogoField,
		addIndustryGroupField,
		skipStage,
		limitStage,
		project,
	})

	if err != nil {
		return nil, err
	}

	err = aggregateVal.All(ctx, &items)

	if err != nil {
		return nil, err
	}

	return items, nil
}
