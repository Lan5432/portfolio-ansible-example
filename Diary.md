# Making the containers work via SSH
- Thought I had to copy my public ssh key to the server but what you need to copy is the key's contents to authorized_keys to the user folder.
    - To add a key to a server you either use the appropriate command or append the SSH key contents to authorized_keys, in the case of Docker that can be done just copying the file, since we know nothing else is on there.
    - Also of course the chmod 700 and chown.
- Didn't know that ssh uses /home/USER/.ssh/id_rsa file by default, I had two keys, one of them had nothing to do with this exercise and that's the one being used. I had to specify the -i flag for the file
- Tried looking for a way to put both containers in different IPs, but ended up just moving the SSH port from 22 to different ones and using a custom inventory file. 
    - I had to specify the port, and use the docker network IP which was 127.0.1.1
- I was confusing become_user and remote_user. remote_user is the user Ansible will use to ssh to the server and become_user is the user Ansible will switch to and run tasks while on the server.
- For introducing stuff in the hosts file I changed to a Docker Compose file and did it that way.

# Proxy
- I was getting too hung up on paths and IP resolution that I didn't realize that in the most basic example of a proxy (I assume they can get more complicated), you're meant to bind the hosts you proxy to, to the path of the host so you access them via its IP.
- 

# 