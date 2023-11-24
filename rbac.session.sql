-- CREATE TABLE clients (
--     id varchar(64) PRIMARY KEY,
--     secret VARCHAR(64) NOT NULL,
--     role VARCHAR(50) NOT NULL,
--     redirect_uri text NOT NULL
-- );
-- INSERT INTO clients (id, secret, role, redirect_uri)
-- VALUES (
--         "laama",
--         "secretkeythis",
--         "editor",
--         "localhost:3000/auth/callback"
--     );
-- CREATE TABLE authorization_codes (
--     code VARCHAR(64) PRIMARY KEY,
--     client_id VARCHAR(64) NOT NULL,
--     user_id VARCHAR(64) NOT NULL,
--     code_challenge VARCHAR(64) NOT NULL,
--     FOREIGN KEY (client_id) REFERENCES clients(id),
--     FOREIGN KEY (user_id) REFERENCES users(id)
-- );
-- CREATE TABLE access_tokens (
--     token VARCHAR(64) PRIMARY KEY,
--     client_id VARCHAR(64) NOT NULL,
--     user_id VARCHAR(64) NOT NULL,
--     FOREIGN KEY (client_id) REFERENCES clients(id),
--     FOREIGN KEY (user_id) REFERENCES users(id)
-- );
SELECT client_id,
    role,
    user_id,
    code_challenge
FROM authorization_codes
    inner join clients on authorization_codes.client_id = clients.id
WHERE code = "znPAJwnwGFBpUmG6zKME_G2iz3mO15Ru5Wls71TtXzw=";