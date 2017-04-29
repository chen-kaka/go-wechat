
rsync -cavzP --delete-after ./ --exclude-from='.rsync-exclude' root@123.207.1.214:/xy/src/go/src/go-wechat
ssh root@123.207.1.214 "\
  cd /xy/src/go/src/go-wechat; \
  sh run.sh; \
  exit; \
  "