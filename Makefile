ecr = public.ecr.aws/w6n7a8r1
region = us-east-1

publish:
	KO_DOCKER_REPO=$(ecr)/ecs-metadata-proxy ko publish --bare ./

login:
	aws ecr-public get-login-password --region $(region) | docker login --username AWS --password-stdin $(ecr)
