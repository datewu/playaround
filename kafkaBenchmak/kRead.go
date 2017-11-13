package main

import (
	"github.com/Shopify/sarama"

	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	addr    = flag.String("addr", ":8080", "The address to bind to")
	brokers = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
	verbose = flag.Bool("verbose", false, "Turn on Sarama logging")
)

const topic = "important"

func (s *Server) collectQueryStringData() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		consumer, err := s.Master.ConsumePartition(topic, 0, sarama.OffsetOldest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to init consumer:, %s", err)
		}
		select {
		case err := <-consumer.Errors():
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "consume err:, %s", err)
		case msg := <-consumer.Messages():
			fmt.Fprintf(w, "consume err:, %s", err)
			fmt.Fprintf(w, "Received messages", string(msg.Key), string(msg.Value))
		}
	})
}

func main() {
	flag.Parse()

	if *verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	server := &Server{
		Master: newDataCollector(brokerList),
	}
	defer func() {
		if err := server.Close(); err != nil {
			log.Println("Failed to close server", err)
		}
	}()

	log.Fatal(server.Run(*addr))
}

type Server struct {
	Master sarama.Consumer
}

func (s *Server) Close() error {
	if err := s.Master.Close(); err != nil {
		log.Println("Failed to shut down data collector cleanly", err)
	}

	return nil
}

func (s *Server) Handler() http.Handler {
	return s.collectQueryStringData()
}

func (s *Server) Run(addr string) error {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.Handler(),
	}

	log.Printf("Listening for requests on %s...\n", addr)
	return httpServer.ListenAndServe()
}
func newDataCollector(brokerList []string) sarama.Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	return master
}
