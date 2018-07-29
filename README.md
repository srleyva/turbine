# turbine
Overview: 
The purpose of this project is to help businesses manage their internal workflows and track data in something as a call-center.

## Goal when it comes to infra
https://www.youtube.com/watch?v=kOa_llowQ1c

## Design:
* Go
*  Microservice Architecture 
	* While it tried to remain agnostic to the infrastructure it is designed with kubernetes in mine
	*  [mico](https://github.com/micro/micro) chosen as the framework because of being opinionated and quick to get off the ground
		* micro cli is helpful when troubleshooting services
		* service discovery is simple and supports multiple service registries (consul, etcd, k8s, etc.)
 * Backend For Frontend 
 * Angular/Bootstrap
 	* opinionated

## TODOs

- [ ] - Pipeline!

## Core-Services
- [x] - User-service
	- [x] - Managing users and employees
	- [ ] - Managing rights

- [x] - authentication-service 
	- [x] - Manage verifying user is who they say they are
	- Support 
		- [x] - BasicAuth
		- [ ] - LDAP
		- [ ] - oAuth
			- [ ] - Google
			- [ ] - Facebook

- [ ] - task-service
	- [ ] - manage and track tasks and related data 
	- [ ] - support heisenberg matrix and kanban style of working

- [ ] - call-data (python)
	- [ ] - track call data
	- [ ] - calculate metrics using ML and expose metrics

## Frontend Features 

- [ ] - KanBan Task Tracker
- [ ] - CallCenter Logger and Metrics Viewer
- [ ] - Dashboard (announcements, reminders, etc.)
