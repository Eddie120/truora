DROP TABLE IF EXISTS m_Keys; CREATE TABLE m_Keys(id serial PRIMARY KEY,name VARCHAR(255) NOT NULL,publicKey TEXT NOT NULL,privateKey TEXT NOT NULL);
CREATE INDEX  idx_m_keys_name ON m_keys(name);