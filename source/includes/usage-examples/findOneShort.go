	// begin findOneUsageExample
	coll := client.Database("sample_mflix").Collection("movies")
	// result is an empty bson struct that will be populated with the
	// matched document when the FindOne method returns
	var result bson.D
	err = coll.FindOne(ctx, bson.D{{"title", "The Room"}}).Decode(&result)
	// end findOneUsageExample
