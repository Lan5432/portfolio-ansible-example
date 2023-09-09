Role Name
=========

Create apps in a server meant to be proxied to.

Requirements
------------



Role Variables
--------------
- python_app_port: Port to configure the Python app's uWSGI server on.
- golang_app_port: Port to configure the Go app on.

Dependencies
------------


Example Playbook
----------------

    - hosts: servers
      roles:
         - role: app_server
           vars:
              python_app_port: 9091
              golang_app_port: 9092

License
-------

BSD

Author Information
------------------

 - https://github.com/Lan5432
