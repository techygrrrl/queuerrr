CREATE TABLE queue (
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	twitch_user_id text PRIMARY KEY,
	twitch_username text,
	notes text
);
