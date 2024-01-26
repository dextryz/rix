package main

import (
	"fmt"
	"net/http"

	"github.com/fiatjaf/eventstore/elasticsearch"
	"github.com/fiatjaf/khatru"
)

func main() {

	relay := khatru.NewRelay()

	relay.Info.Name = "Ixian Relay"
	relay.Info.PubKey = "aa327517150b10655711a01708c18753b0a33b1b44d058ac41d368ab15b95190"
	relay.Info.Description = "Ixian relay for fast article content searching"

	db := elasticsearch.ElasticsearchStorage{URL: ""}
	if err := db.Init(); err != nil {
		panic(err)
	}

	relay.StoreEvent = append(relay.StoreEvent, db.SaveEvent)
	relay.QueryEvents = append(relay.QueryEvents, db.QueryEvents)
	relay.CountEvents = append(relay.CountEvents, db.CountEvents)
	relay.DeleteEvent = append(relay.DeleteEvent, db.DeleteEvent)

	fmt.Println("running on :3334")
	http.ListenAndServe(":3334", relay)
}
