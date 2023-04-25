CREATE TABLE IF NOT EXISTS schools(
    id serial not null unique, 
    class_count integer not null,
    name varchar(255) not null unique, 
    password varchar(255) not null
);

CREATE TABLE IF NOT EXISTS admins(
    id serial not null unique, 
    username varchar(255) not null unique, 
    password varchar(255) not null
);

CREATE TABLE IF NOT EXISTS teachers (
    id serial not null unique, 
    school_id integer not null,
    first_name varchar(50) not null, 
    second_name varchar(50) not null, 
    username varchar(50) not null unique, 
    password varchar(255) not null,
    private boolean not null,
    foreign key (school_id) references schools (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS classes (
    id serial not null unique, 
    school_id integer not null,
    teacher_id integer not null,
    grade varchar(20) not null,
    name varchar(20) not null,
    foreign key (school_id) references schools (id) on delete cascade,
    foreign key (teacher_id) references teachers (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS students (
    id serial not null unique, 
    class_id integer not null,
    points integer default 0,
    first_name varchar(255) not null, 
    second_name varchar(255) not null, 
    username varchar(255) not null unique, 
    password varchar(255) not null,
    foreign key (class_id) references classes (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS books (
    id serial not null unique,
    name varchar(255) not null,
    category varchar(255) not null,
    class varchar(255) not null,
    hashed_id varchar(255) not null,
    language varchar(255) not null
);

CREATE TABLE IF NOT EXISTS complete_books (
    id serial not null unique,
    book_id integer not null,
    student_id integer not null,
    points integer not null,
    foreign key (book_id) references books (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS complete_audios (
    id serial not null unique,
    book_id integer not null,
    student_id integer not null,
    points integer not null,
    foreign key (book_id) references books (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS tests (
    id serial unique,
    book_id integer not null,
    lang varchar(25) not null,
    foreign key (book_id) references books (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS questions (
    id serial unique,
    test_id integer not null,
    with_image boolean not null,
    image text,
    audio text,
    question varchar(255) not null,
    foreign key (test_id) references tests (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS answers (
    id serial unique,
    question_id integer not null,
    with_image boolean not null,
    image text,
    audio text,
    answer varchar(255) not null,
    correct boolean not null,
    foreign key (question_id) references questions (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS complete_tests (
    id serial unique,
    test_id integer not null,
    student_id integer not null,
    points integer not null,
    foreign key (test_id) references tests (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS complete_answers (
    id serial unique,
    complete_test_id integer not null,
    question_id integer not null,
    answer_id integer not null,
    foreign key (complete_test_id) references complete_tests (id) on delete cascade,
    foreign key (question_id) references questions (id) on delete cascade,
    foreign key (answer_id) references answers (id) on delete cascade
);



CREATE TABLE IF NOT EXISTS body (
    id serial unique,
    part text not null,
    name text not null unique,
    img_url text not null,
    img_icon_url text not null,
    price integer not null
);

INSERT INTO body(part, name, img_url, img_icon_url, price)
VALUES
    ('head', 'default head', 'http://0.0.0.0:8080/static/default/head.png', 'http://0.0.0.0:8080/static/default/head-icon.png', 0),
    ('chest', 'default chest', 'http://0.0.0.0:8080/static/default/chest.png', 'http://0.0.0.0:8080/static/default/chest-icon.png', 0),
    ('legs', 'default legs', 'http://0.0.0.0:8080/static/default/legs.png', 'http://0.0.0.0:8080/static/default/legs-icon.png', 0),
    ('arms', 'default arms', 'http://0.0.0.0:8080/static/default/arms.png', 'http://0.0.0.0:8080/static/default/arms-icon.png', 0);

CREATE TABLE IF NOT EXISTS current_body ( 
    id serial unique,
    head_id integer not null,
    chest_id integer not null,
    legs_id integer not null,
    arms_id integer not null,
    student_id integer not null,
    foreign key (student_id) references students (id) on delete cascade,
    foreign key (head_id) references body (id) on delete cascade,
    foreign key (chest_id) references body (id) on delete cascade,
    foreign key (legs_id) references body (id) on delete cascade,
    foreign key (arms_id) references body (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS buyed ( 
    id serial unique,
    student_id integer not null,
    body_id integer not null,
    foreign key (student_id) references students (id) on delete cascade,
    foreign key (body_id) references body (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS open_questions (
    id serial unique,
    category text not null,
    question text not null,
    audio_link text
);

CREATE TABLE IF NOT EXISTS open_answers(
    id serial unique,
    answer text not null,
    open_questions_id integer not null,
    student_id integer not null,
    book_id integer not null,
    foreign key (open_questions_id) references open_questions (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade,
    foreign key (book_id) references books (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS open_comments ( 
    id serial unique,
    answer_id integer not null unique,
    teacher_id integer not null,
    student_id integer not null,
    comment text not null,
    audio_link text,
    points integer not null,
    foreign key (answer_id) references open_answers (id) on delete cascade,
    foreign key (teacher_id) references teachers (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS open_notifications (
    id serial unique,
    comments_id integer not null,
    is_read boolean not null,
    foreign key (comments_id) references open_comments(id) on delete cascade
);

CREATE TABLE IF NOT EXISTS creative_questions (
    id serial unique,
    category text not null,
    question text not null,
    audio_link text
);

CREATE TABLE IF NOT EXISTS creative_answers(
    id serial unique,
    answer text not null,
    creative_questions_id integer not null,
    student_id integer not null,
    book_id integer not null,
    foreign key (creative_questions_id) references creative_questions (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade,
    foreign key (book_id) references books (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS creative_comments ( 
    id serial unique,
    answer_id integer not null unique,
    teacher_id integer not null,
    student_id integer not null,
    comment text not null,
    audio_link text,
    points integer not null,
    foreign key (answer_id) references creative_answers (id) on delete cascade,
    foreign key (teacher_id) references teachers (id) on delete cascade,
    foreign key (student_id) references students (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS creative_notifications (
    id serial unique,
    comments_id integer not null,
    is_read boolean not null,
    foreign key (comments_id) references creative_comments(id) on delete cascade
);