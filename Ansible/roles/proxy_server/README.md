Role Name
=========

Simple role to create an proxy_server server with its own page and a proxy to two apps, a Python and a Golang one.

Requirements
------------



Role Variables
--------------

 - python_app_ip: Host of Python app
 - golang_app_ip: Host of Go app
 - python_app_port: Port of Python app in its host
 - golang_app_port: Port of Go app in its host

Dependencies
------------


Example Playbook
----------------
    - hosts: servers
      roles:
         - role: proxy_server
            vars:
                python_app_ip: python.app
                golang_app_ip: go.app
                python_app_port: 9090
                golang_app_port: 9090

License
-------

BSD

Author Information
------------------

 - https://github.com/Lan5432
