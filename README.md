# gcp-app-engine-go

Experimentation for backend on GCP's App Engine
Archived.

App Engine project with multiples services and CI/CD setup:
- default (hello)
- health
- user

* Github Actions:
- Build on develop commit
- Build & Deploy on main commit

# To open the running app:
gcloud app browse

# To deploy
cd {serviceFolder}
gcloud app deploy

# To update routing
Update dispatch.yaml
gcloud app deploy dispatch.yaml

# To generate the deployment keys for an account with deploy permissions:
gcloud iam service-accounts keys create sa-swiftshop-key \
    --iam-account=s{projectname}@appspot.gserviceaccount.com
