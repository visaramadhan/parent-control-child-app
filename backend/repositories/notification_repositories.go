package notification

import (
    "context"
    "firebase.google.com/go/v4/messaging"
)

func SendNotification(fcmToken, title, body string) error {
    ctx := context.Background()
    message := &messaging.Message{
        Token: fcmToken,
        Notification: &messaging.Notification{
            Title: title,
            Body:  body,
        },
    }

    _, err := firebase.MessagingClient.Send(ctx, message)
    return err
}
