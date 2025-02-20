package queue

// func StartQueueServices(cont container.Container) {
// 	fmt.Println("Starting queue service...")
// 	ctx := context.Background()

// 	go func() {
// 		err := cont.QueueServices.ConsumeData(ctx, cont.EnvironmentConfig.RabbitMq.ProducerName)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	}()

// 	go func() {
// 		err := cont.QueueServices.ConsumeData(ctx, cont.EnvironmentConfig.RabbitMq.ConsumerName)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	}()
// }
