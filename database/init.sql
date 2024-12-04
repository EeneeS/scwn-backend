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

create table if not exists publishTargets (
  id int auto_increment primary key,
  project_id char(36),
  platform varchar(255),
  url varchar(255),
  foreign key (project_id) references projects(id)
);
