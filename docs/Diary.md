This diary is mostly for myself as I stop and retake personal projects somewhat randomly depending on the day to day. And also so anyone can see my problem analysis.
# Making the containers work via SSH
- I'm not completely oblivious to SSH as this file might imply, but I just forget how it works from not having to set it up often.
- First thing to mention, my PC is a Windows machine, so to more easily work with Linux style tools I installed the Windows Subsystem for Linux.
    - This means the ssh-agent has to be started for a terminal, and the keys you want to use, added.
    - And I kept forgetting that if you want something done by a script to the environment to stick, you need to use the `source` command.
- Thought I had to copy my public ssh key to the server but what you need to copy is the key's contents to authorized_keys to the user folder.
    - To add a key to a server you either use the appropriate command or append the SSH key contents to authorized_keys, in the case of Docker that can be done just copying the file, since we know nothing else is on there.
    - Also of course the chmod 700 and chown.
    - And I try to make my life a bit easier by creating a 
- Didn't know that ssh uses /home/USER/.ssh/id_rsa file by default, I had two keys, one of them had nothing to do with this exercise and that's the one being used. I had to specify the -i flag for the file
- Tried looking for a way to put both containers in different IPs, but ended up just moving the SSH port from 22 to different ones and using a custom inventory file. 
    - I had to specify the port, and use the docker network IP which was 127.0.1.1
- For introducing stuff in the hosts file I changed to a Docker Compose file and did it that way.

# Ansible
- I was confusing become_user and remote_user. remote_user is the user Ansible will use to ssh to the server and become_user is the user Ansible will switch to and run tasks while on the server.

# Proxy
- I was getting too hung up on paths and IP resolution that I didn't realize that in the most basic example of a proxy (I assume they can get more complicated), you're meant to bind the hosts you proxy to, to the path of the host so you access them via its IP. So if your host is "www.host.com" one of your paths would be "www.host.com/a-proxy".
- Issues with how to get the apps.conf file to answer for the "python" and "golang" paths, because I had an issue of these ending up in the local-page.conf site.
    - One idea I decided to apply was to create a default site to fall into, with a 404, at least as a clear sign that I'm doing something wrong.
- Whole odyssey about why my proxy to Python was failing with `connect() failed (111: Unknown error) while connecting to upstream`, the things I misunderstood:
    - I had misunderstood how uWSGI was working, I had put `http-socket` which means the server rises and talks HTTP through a port, as opposed to regular uWSGI. But then I had `uwsgi_pass` in my Nginx proxy, so I was talking uwsgi.
    - I was literally telling container A to talk to container B via "localhost", which makes no sense but I did not give it the proper thought. I then took advantage of the docker-compose networking to give container B's IP to container A via name resolution.
     

# 

[Back to README.md](../README.md)