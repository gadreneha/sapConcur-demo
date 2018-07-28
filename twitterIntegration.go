ticker := time.NewTicker(time.Second * 1).C
go func() {
    for {
        select {
        case <- ticker:
            response,_ := http.Get("http://...")
            _, err := io.Copy(os.Stdout, response.Body)
            if err != nil {
                log.Fatal(err)
            }
            response.Body.Close()
        }
    }

}()


time.Sleep(time.Second * 10)
