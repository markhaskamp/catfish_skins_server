 
## Tasks

### Doing

- query sqlite database from handleLogin 

### To Do

- main page shows all scores
- admin page for id/name/pwd table
- on successful login, write id as cookie

### Done



### Auth Notes

- server storage has:
  - id, name, pwd
- server storage is:
  - environment variables ?
  - sqlite ?
- login:
  - input name/pwd
  - process queries for combo
  - finds it:
    - returns id
    - id is stored as cookie
  - doesn't find it:
    - returns 401


### Start server on boot

edit _/etc/rc.d/rc.local_  
add _/home/ec2-user/service/catfish\_skins\_server/bin/server_



