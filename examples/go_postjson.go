m := map[string]interface{}{
	"name":    "backy",
	"species": "dog",
}
mJson, err := json.Marshal(m)
if err != nil {
	panic(err)
}
contentReader := bytes.NewReader(mJson)
req, err := http.NewRequest("POST", "http://example.com", contentReader)
if err != nil {
	panic(err)
}
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Notes", "GoRequest is coming!")
client := &http.Client{}
resp, err := client.Do(req)
