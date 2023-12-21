CREATE TABLE followers (
  id BIGSERIAL PRIMARY KEY,

  user_id BIGINT NOT NULL,
  following_user_id BIGINT NOT NULL,

  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE
);

ALTER TABLE followers ADD CONSTRAINT fk_followers_user_id FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE followers ADD CONSTRAINT fk_followers_following_user_id FOREIGN KEY (following_user_id) REFERENCES users(id);
