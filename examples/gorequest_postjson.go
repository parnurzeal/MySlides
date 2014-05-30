request := gorequest.New()
resp, body, errs := request.Post("http://example.com").
	Set("Notes", "gorequst is coming!").
	Send(`{"name":"backy", "species":"dog"}`).
	End()
