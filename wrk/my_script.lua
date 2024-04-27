request = function()
    wrk.method = "GET"
    wrk.headers["Content-Type"] = "application/json"
    wrk.body = '{"key":"value"}'
    return wrk.format()
   end
   