package Topic

import (
    "io/ioutil"
    "net/http"
    "time"
    "fmt"
)


type TopicClient struct {
    Token  string
}

func (client TopicClient) ServerCountPost(servers string)  string {

    req, err := http.NewRequest("POST", "https://api.topiclist/bots/stats",  nil)

    if err != nil {
        return ""
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", client.Token)
    req.Header.Set("serverCount", servers)

    r, err := httpClient.Do(req)


    if err != nil {
        return "Bro There is  Error"
    }

      defer r.Body.Close()

    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return "Bro There is  Error"
    }

    return string(body)



}

func (client TopicClient) HasVoted(id string) string {

      req, err := http.NewRequest("GET", fmt.Sprintf("https://api.topiclist.xyz/bots/check/", id),  nil)

    if err != nil {
        return "Bro There is  Error"
    }

    req.Header.Set("Authorization", client.Token)
    req.Header.Set("Content-Type", "application/json")

    r, err := httpClient.Do(req)


      if err != nil {
            return "Bro There is  Error"
      }

      defer r.Body.Close()


    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return "Bro There is  Error"
    }

    return string(body)

}

func (client TopicClient) Search(id string) string {

    req, err := http.NewRequest("GET", fmt.Sprintf("https://api.topiclist.xyz/bots/", id),  nil)

    if err != nil {
        return "Bro There is  Error"
    }

    req.Header.Set("Content-Type", "application/json")

    r, err := httpClient.Do(req)


    if err != nil {
            return "Bro There is  Error"
      }

      defer r.Body.Close()

    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return "Bro There is  Error"
    }

    return string(body)

}



var httpClient = &http.Client{Timeout: 10 * time.Second}