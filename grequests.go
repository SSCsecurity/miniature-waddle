#####

######





resp, err := grequests.Patch("http://api.stitchfix.com/style?type=jeans",
            &grequests.RequestOptions{
                Headers: getHeaders(), 
                JSON:   map[string]string{"token": req.Header.Get("X-Riscosity-Tkn"), "ssn": "222-22-2222"},

                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })








########

resp, err := grequests.Post("https://api.twilio.com/pants?type=jeans",
            &grequests.RequestOptions{
                Headers: getHeaders(), 
                JSON:   map[string]string{"token": req.Header.Get("X-Riscosity-Tkn"), "ssn": "111-11-1111"},

                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })


###
resp, err := grequests.Post("http://api.sendgrid.com/pants?type=jeans",
           getRequestOptions())
####
####
####

resp, err := grequests.Post(getURL(),
            &grequests.RequestOptions{
                Headers: getHeaders(), 
JSON:   map[string]string{"token": req.Header.Get("X-Riscosity-Tkn")},

                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })

resp, err := grequests.Get("https://oldnavy.com/shirts",
            &grequests.RequestOptions{
                JSON:   map[string]string{"token": req.Header.Get("X-Riscosity-Tkn"), "phone_number": "555-555-5555"},
                Headers: map[string]string{"token": req.Header.Get("X-Riscosity-Tkn")},

                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })



resp, err := grequests.Post("http://api.okta.com/pants?type=jorts",
            &grequests.RequestOptions{Headers: getHeaders()})

resp, err := grequests.Post("http://api.slack.com/pants?type=jeans",
            &grequests.RequestOptions{
                Headers: getHeaders(test,variable)})

resp, err := grequests.Post("http://extapi.databricks.com/pants?type=jeans",
            &grequests.RequestOptions{
                Headers: headers,})


resp, err := grequests.Get(https://api.gap.com/shirts", _)




url := "https://utilityapi.com/monkey"
resp, err := grequests.Get(url,
            &grequests.RequestOptions{
                JSON:   map[string]string{"toke2": req.Header.Get("X-Riscosity-Tkn")},
                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })


resp, err := grequests.Get(getURL(),
            &grequests.RequestOptions{
                JSON:   map[string]string{"token4": req.Header.Get("X-Riscosity-Tkn")},
                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })

url = "https://qpi.net/cheese"




######
#####
####

resp, err := grequests.Patch("https://paymentgateway.plaid.com/clothes?type=sweaters",
            &grequests.RequestOptions{
                JSON:   map[string]string{"token": req.Header.Get("X-Riscosity-Tkn"), "test": "1234"},
                IsAjax: true, // this adds the X-Requested-With: XMLHttpRequest header
            })

resp, err := grequests.Delete("http://microsoft.com/pants?type=jeans",
            &grequests.RequestOptions{JSON:   getJSON()}
)
