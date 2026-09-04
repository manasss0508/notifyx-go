-- +goose Up
CREATE TABLE IF NOT EXISTS templates (
                                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    channel TEXT NOT NULL,
    subject TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

INSERT INTO templates (name, channel, subject, body) VALUES
(
 'welcome',
 'MAIL',
 'Welcome to NotifyX, {{name}}!',
 'Hello {{name}},
 Welcome to NotifyX! We''re excited to have you on board.
 Start exploring and enjoy your experience.
 Thanks,
 The NotifyX Team'
),
(
 'otp',
 'MAIL',
 'Your OTP Code',
 'Hello {{name}},
 Your One-Time Password is:
 {{otp}}
 This code will expire in {{expiry_minutes}} minutes.
 If you didn''t request this code, please ignore this MAIL.'
),

(
 'password_reset',
 'MAIL',
 'Reset Your Password',
 'Hello {{name}},
 We received a request to reset your password.
 Click the link below to continue:
 {{reset_link}}
 This link expires in {{expiry_minutes}} minutes.'
),

(
 'MAIL_verification',
 'MAIL',
 'Verify Your MAIL Address',
 'Hello {{name}},
 Please verify your MAIL address by clicking the link below.
 {{verification_link}}
 Thanks for joining NotifyX.'
),

(
 'login_alert',
 'MAIL',
 'New Login Detected',
 'Hello {{name}},
 A new login was detected on your account.
 Time: {{login_time}}
 Location: {{location}}
 Device: {{device}}
 If this wasn''t you, please secure your account immediately.'
),

(
 'order_confirmation',
 'MAIL',
 'Order {{order_id}} Confirmed',
 'Hello {{name}},
 Your order has been confirmed.
 Order ID: {{order_id}}
 Amount: {{amount}}
 Currency: {{currency}}
 We''ll notify you when it ships.'
),

(
 'order_shipped',
 'MAIL',
 'Your Order Is On The Way',
 'Hello {{name}},
 Great news!
 Your order {{order_id}} has been shipped.
 Tracking Number: {{tracking_number}}
 Estimated Delivery: {{estimated_delivery}}'
),

(
 'payment_success',
 'MAIL',
 'Payment Successful',
 'Hello {{name}},
 Your payment was successfully processed.
 Transaction ID: {{transaction_id}}
     Amount: {{amount}}
     Currency: {{currency}}
 Thank you for your payment.'
),

(
 'payment_failed',
 'MAIL',
 'Payment Failed',
 'Hello {{name}},
 Unfortunately, your recent payment couldn''t be completed.
 Amount: {{amount}}
     Currency: {{currency}}
     Reason: {{reason}}
 Please try again or use another payment method.'
),

(
 'subscription_renewal',
 'MAIL',
 'Subscription Renewed',
 'Hello {{name}},
 Your {{plan_name}} subscription has been successfully renewed.
 Renewal Date: {{renewal_date}}
 Next Billing Date: {{next_billing_date}}
 Thank you for staying with NotifyX.'
);

-- +goose Down
DROP TABLE IF EXISTS templates;

