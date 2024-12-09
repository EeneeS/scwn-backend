create table if not exists users (
  id char(36) PRIMARY KEY,
  email varchar(255),
  created_at datetime
);

create table if not exists projects (
  id char(36) primary key,
  name varchar(255),
  user_id char(36),
  foreign key (user_id) references users(id)
);

create table if not exists publish_targets (
  id int auto_increment primary key,
  project_id char(36),
  platform varchar(255),
  url varchar(255),
  foreign key (project_id) references projects(id)
);

create table if not exists changes (
  id int auto_increment primary key,
  project_id char(36),
  element varchar(255),
  type varchar(255),
  original_value varchar(255),
  new_value varchar(255),
  route varchar(255),
  foreign key (project_id) references projects(id)
)
