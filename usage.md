## Detailed instructions for use

https://www.wolai.com/ctfhub/3DvnJJtPbHyyDtVkDaW1yz

## Require

All environments are built based on amd64

## Create a question

Please follow the following process to create

0. Configure producer information

Producer information only needs to be configured once. If the configuration is incorrect, please use the `cg config clean` command to clear and reconfigure it.

```bash
cg config set
```

1. Use the wizard to generate question templates

Created using the `new` subcommand
```bash
cg new wizard
```

If you have special requirements for the flag placement location, please select `Yes` for the `Separately process flag location` option in the wizard, and modify the `environment/files/flag.sh` file after the creation is complete.

If you have special requirements for the service startup of the environment, please select `Yes` in the `Process part of the service startup separately` option in the wizard, and modify the `environment/files/start.sh` file after the creation is completed.

2. Complete the image by referring to the example

Please make sure that the base image you are using is the latest. You can pull the corresponding base image before building.

If there are no special needs after generation, please try not to change the dockerfile as much as possible

Example link: https://github.com/ctfhub-team/base_image/

3. Test image

Using `docker` subcommand for environment testing
```bash
cg docker auto
```

4. Modify meta information

Edit `meta.yaml` in the question directory. Please refer to `meta description` for the specific meaning.

5. Packing

Use zip to package the entire question folder

```bash
zip -r xxx.zip xxx
```

6. Upload

Please send the packaged question folder to the relevant person in charge

## meta description

```yaml
author:
   # Producer ID, automatically generated by cg
   name: l1n3
   # Producer email, automatically generated by cg
   contact: yw9381@163.com
task:
   # Question image name, automatically generated by cg
   name: challenge_web_2022_hitcon_rce
   # Question type, automatically generated by cg
   type: Web
   #Problem description
   description: aasasdsadas
   # Question difficulty, automatically generated by cg
   level: sign in
   # topic flag
   # If it is a static flag, fill in the specific flag value here
   # If it is a dynamic flag, this is an empty string.
   flag:
   #Tips prompt, if there is no prompt, this will be an empty array
   hints:
     - asdasd
     -asdas
     -asdad
challenge:
   # Question display name
   name: rce
   # Question source, source format: year - abbreviation of competition name - question type - question display name
   # For example, the Web class babysqli of the 2021 QWB Cup is 2021-QWB-Web-baysqli
   # For example, the babyheap of the Pwn class of SCTF in 2019 is 2019-SCTF-Pwn-bayheap
   refer: 2022-HitCon-Web-rce
   #Title label. The label should reflect the test points of the question as much as possible. If there is no label, this will be an empty array.
   tags:
     - web
     - 2022
     - hitcon

```
## Known issues

1. The Python/NodeJS/Ruby/Java basic environment has not been completed yet.
2. The check function has not been completed yet. If the production is completed, please directly send the packaged file to the docking person, who will be responsible for review and testing.