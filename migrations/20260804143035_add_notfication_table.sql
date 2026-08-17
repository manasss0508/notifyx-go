-- +goose Up
CREATE TABLE notifications
(
    id             UUID PRIMARY KEY NOT NULL,           -- notification id
    channel TEXT NOT NULL ,
    recipient      TEXT             NOT NULL,           -- mail or number
    template       TEXT             NOT NULL,           -- message template
    variables           TEXT             NOT NULL,           -- variables of template
    status         TEXT             NOT NULL,           -- status of notification "PENDING", "PROCESSING", "SENT", "FAILED", "RETRYING"
    priority       TEXT             NOT NULL,           -- "LOW", "MEDIUM", "HIGH"
    scheduled_at   TIMESTAMP WITH TIME ZONE,                           -- when notification should be sent
    retry_count    INT     NOT NULL  DEFAULT 0,                 -- count of retry sending
    max_retry      INT     NOT NULL  DEFAULT 3,                 -- max count of retry sending
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP, -- notfication created at
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP, -- notifcation updated at
    sent_at        TIMESTAMP WITH TIME ZONE,                           -- notification sent at
    failure_reason TEXT                                 -- reason why message was failed to sent
);

-- +goose Down
DROP TABLE notifications;
