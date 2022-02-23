client := &http.Client{
        Timeout: time.Second * 10,
 }

req, err := http.NewRequest(method,url, nil)
response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()

////



req, err := http.NewRequest(http.MethodPatch, url(), body)
response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()


req, err := http.NewRequest("
GET", url(), body)
response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

 defer response.Body.Close()




req, err := http.NewRequest('Delete', "https://microsoft.com/cookies?=jarvis", body)
response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()




req, err := http.NewRequest('PUT', "https://microsoft.com/cookies?=jarvis", bytes.NewBuffer(payload)
)



response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()

req, err := http.NewRequest('POST', "https://microsoft.com/cookies?=jarvis", bytes.NewBuffer(payload))


response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()





req, err := http.NewRequest('Delete', "https://microsoft.com/cookies?=jarvis", bytes.NewBuffer(payload))


response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()


req, err := http.NewRequest('PATCH', "https://microsoft.com/cookies?=jarvis", bytes.NewBuffer(payload))  \\test

response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()

//
req, err := http.NewRequest(
'patch', "https://microsoft.com/cookies?=jarvis",bytes.NewBuffer({
})
) 
if err != nil {
}
//////
response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()
/////
req, err := http.NewRequest(method, url, nil)
if err != nil {
  return fmt.Errorf("Got error %s", err.Error())
}
    req.Header.Set("user-agent", "golang application")
   req.Header.Add("foo", 
"bar1")  
   req.Header.Add("foo", getValue(), 
) 

response, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("Got error %s", err.Error())
    }

   defer response.Body.Close()

