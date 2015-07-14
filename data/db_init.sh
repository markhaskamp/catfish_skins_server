#!/bin/sh

sqlite3 catfish.db 'drop table players;'
sqlite3 catfish.db 'CREATE TABLE players (id integer, name text);'
sqlite3 catfish.db 'INSERT INTO "players" VALUES(1,"Alan");'
sqlite3 catfish.db 'INSERT INTO "players" VALUES(2,"Dave");'
sqlite3 catfish.db 'INSERT INTO "players" VALUES(3,"Doug");'
sqlite3 catfish.db 'INSERT INTO "players" VALUES(4,"Keith");'
sqlite3 catfish.db 'INSERT INTO "players" VALUES(5,"Mark");'
sqlite3 catfish.db 'INSERT INTO "players" VALUES(6,"Ric");'
sqlite3 catfish.db 'select * from players;'

