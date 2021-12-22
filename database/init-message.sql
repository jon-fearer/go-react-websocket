CREATE TABLE message
(
    message_id   SERIAL PRIMARY KEY,
    message_text TEXT NOT NULL
);

CREATE FUNCTION message_notify()
    RETURNS trigger AS
$$
BEGIN
	PERFORM pg_notify('message_channel', NEW.message_text);
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER message
    AFTER INSERT ON message
    FOR EACH ROW
    EXECUTE PROCEDURE message_notify();
