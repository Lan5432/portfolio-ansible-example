Portfolio Ansible Example
=====================

Example of Ansible playbook and roles that I did to remember and expand my knowledge of the tool that I gained from university.

It features two roles, I realize that the apps one could've been split in two separate machines, but left it this way for simplicity's sake.
 - proxy_server: a proxy server that redirects to the separate apps.
 - app_server: the target host that holds both a Python Flask app and a Golang one.


Support resources
---------------------
 - On top of my own research on how to do things properly, I have to thank Jose Manuel Redondo for lending me his teaching materials on the subject. https://www.researchgate.net/profile/Jose-Redondo-3
 - Used these config files as a basis for the proxy_server role: https://github.com/h5bp/server-configs-proxy_server


Documentation
---------------------
- [My diary of how I went through solving my issues](./docs/Diary.md)
- [Setup](./docs/Setup.md)