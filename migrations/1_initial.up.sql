CREATE TABLE IF NOT EXISTS projects
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO projects(name) VALUES ('Project X');

CREATE TABLE IF NOT EXISTS goods
(
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    priority INT NOT NULL,
    removed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE goods ADD FOREIGN KEY(project_id) REFERENCES projects(id);

CREATE OR REPLACE FUNCTION increment_priority()
RETURNS TRIGGER AS $$
BEGIN
    NEW.priority := COALESCE((SELECT MAX(priority) FROM goods), 0) + 1;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_increment_priority
BEFORE INSERT ON goods
FOR EACH ROW
EXECUTE FUNCTION increment_priority();


CREATE OR REPLACE FUNCTION update_priority(good_id INT, proj_id INT, new_priority INT)
    RETURNS TABLE(updated_id INT, updated_priority INT) AS $$
DECLARE row RECORD;
BEGIN
    UPDATE goods SET priority=new_priority WHERE id=good_id AND project_id=proj_id;
    RETURN QUERY SELECT good_id, new_priority;
    FOR row IN SELECT id FROM goods WHERE id <> good_id AND removed=false ORDER BY priority LOOP
            IF new_priority <= (SELECT priority FROM goods WHERE id=row.id) THEN
                new_priority := new_priority + 1;
                UPDATE goods SET priority=new_priority WHERE id=row.id;
                RETURN QUERY SELECT row.id, new_priority;
            END IF;
        END LOOP;
END;
$$ LANGUAGE plpgsql;
