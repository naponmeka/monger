for SIZE in 16 32 64 128 256 512 1024
do
    sips -z $SIZE $SIZE icon.png --out icon_${SIZE}x${SIZE}.png
done
