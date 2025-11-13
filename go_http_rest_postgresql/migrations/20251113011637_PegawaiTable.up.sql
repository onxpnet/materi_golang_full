-- Add up migration script here
create table if not exists pegawai (
  id serial primary key,
  nama text not null,
  jabatan text not null
);
