CREATE TABLE to_do_task(
                           user_id uuid,
                           CONSTRAINT fk_user_id
                               FOREIGN KEY (user_id) REFERENCES user_profile(id),
                           task_title VARCHAR,
                           task_description VARCHAR,
                           task_time VARCHAR

)