git-ssh:
	eval "$$(ssh-agent -s)"
	ssh-add ~/.ssh/newgit_id_rsa

.PHONY: git-ssh