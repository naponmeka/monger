qtdeploy build darwin .
cd deploy/darwin
mv monger.app Monger.app
create-dmg Monger.app
mv Monger\ 1.0.0.dmg Monger1.0.0-MacOS-alpha.dmg
cp Monger1.0.0-MacOS-alpha.dmg ../../releases