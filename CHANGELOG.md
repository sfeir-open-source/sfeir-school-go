# Changelog

## v1.0.0 
- 19/10/08 chore(dep): move to go modules for master branch

## v0.0.4
- 18/03/10 feat(dao): add postgresql and db migration support for dao
           fix(make): add support for windows in Makefile

## v0.0.3
- 18/02/19 chore(build): update Docker build images to golang v1.10
           feat(web): add CORS middleware

## v0.0.2
- 17/11/03 fix(test): fix time comparison error in test with go 1.9 (SFR)
           chore(docker): fix docker image and mongo upgrade to 3.4
- 17/10/05 fix(test): enhance web test errors on full server test (SFR)
           refact(model): refact model constructor
           fix(dao): fix error name
           refact(main): update asciiart
- 17/10/04 chore(make): clean make file auto generated help (SFR for RLE)
           chore(dep): update package manager to dep, update vendor
           chore(docker): update docker image version, add multistep build
           refact(main): refact arg parsing, remove redondant arg parsing
- 17/04/14 doc(stat): clarify PlusOne go doc (SFR)
           doc(model): fix the NewID documentation to match method content

## v0.0.1 [17/03/21]
- 17/03/21 fix(main): fix the statd flag help message (SFR)
- 17/03/20 fix(web): fix the Fatal log level in server build (SFR)
- 17/03/17 refact(web): rename handler to controller (SFR)
           chore(test): add postman collection for testing
           chore(doc): update vendor manager in README
           refact(dao): interface type safety check, json omitempty tag
           chore(make): make the bench tool work again
           refact(web): update web test to use MongoDB
           fix(main): fix timezone for bson/json marshalling
- 17/03/14 refact(dao): go fmt src, update DAO with typed const (SFR)
           chore(make): clean the clean target to prevent glide.lock removal
- 17/03/12 refact(dao): update model and dao to use UUID (SFR)
           chore(doc): update main architecture diagram
- 17/02/21 chore(vendor): update vendors to glide (SFR)
           refact(model): add const types
           chore(etc): update scripts and types to todos
- 16/12/20 chore(all): refactor from handsongo to todolost (OFU)

