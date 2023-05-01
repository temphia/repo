
#!/usr/bin/env bash

cd testground

# echo $PWD

# export JAIL=$PWD/jail
# mkdir -p $JAIL
# mkdir -p $JAIL/{bin,lib64}
# cd $JAIL

# cp -v /bin/{bash,ls} $JAIL/bin

# # Get a list of required libraries for bash using ldd command
# LIBRARIES=$(ldd /bin/bash | awk '{print $3}')

# # Copy the required libraries to the $JAIL/lib64 directory
# for LIBRARY in $LIBRARIES; do
#     cp -v "$LIBRARY" $JAIL/lib64
# done

# chmod +x $JAIL/bin/*

# chroot $JAIL /bin/bash

# # chroot: failed to run command ‘/bin/bash’: No such file or directory


# create a temporary directory to use as the guest's filesystem
guest_fs=$(mktemp -d)

# create a directory to use as the bind mount for the guest's filesystem
bind_mount=$(mktemp -d)

# mount the guest's filesystem as a bind mount
sudo mount --bind $guest_fs $bind_mount

# copy all host binaries to the bind mount
sudo cp -r /bin $bind_mount

# chroot into the guest's filesystem
sudo chroot $guest_fs /bin/bash

# unmount the guest's filesystem and remove the temporary directories
sudo umount $bind_mount
rmdir --ignore-fail-on-non-empty $bind_mount
rmdir --ignore-fail-on-non-empty $guest_fs

