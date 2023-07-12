package test

/*func TestGetMessage(t *testing.T) {
	router := initMongoDB()
	apitest.New(). // configuration
			Debug().
			Handler(router).
			Get("/v1/records"). // request
			Expect(t).
			Assert(func(response *http.Response, request *http.Request) error {

			var resBody = make([]byte, len(responses["allRecords"])+1)
			response.Body.Read(resBody)
			if string(resBody) == responses["allRecords"] {
				return nil
			}
			println("expected: ", responses["allRecords"])
			println("got     : ", string(resBody))
			return errors.New("Body not as expected")
		}).
		Status(http.StatusOK).
		End()
}*/
