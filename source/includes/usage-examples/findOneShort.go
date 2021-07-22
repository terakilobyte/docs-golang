	// begin findOneUsageExample
	coll := client.Database("sample_mflix").Collection("movies")
	var result bson.D
	err = coll.FindOne(ctx, bson.D{{"title", "The Room"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}

		log.Panic(err)
	}
	// end findOneUsageExample
