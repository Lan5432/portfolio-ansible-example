Role Name
=========

Simple role to create an Nginx server with its own page and a proxy to two apps, a Python and a Golang one.

Requirements
------------

No requirements.

Role Variables
--------------

 - app_host: The host that will hold the apps being proxied to.

Dependencies
------------


Example Playbook
----------------

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for users too:

    - hosts: servers
      roles:
         - proxy_nginx
           vars:
             app_host: 192.168.56.11

License
-------

BSD

Author Information
------------------

 - https://github.com/Lan5432
