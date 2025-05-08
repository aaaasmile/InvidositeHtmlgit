# Aggiornare il sito
Di solito si crea un nuovo post. Per questo occorre _vido-preproc.exe_. 
Meglio usare start_code.ps1 ed usare il Terminal per fare partire _vido-preproc.exe_.
In Visual Code poi si cambiano tutti gli altri files in src che non sono nell'indice.
_vido-preproc.exe_ viene usato anche per lanciare Webgen.

Ultimamente uso Visual Code per editare tutti i contenuti, compreso l'ultimo post,
in quanto lo Spell Checker funziona molto bene (extension Italian - Code Spell Checker). 
Questi sono i due comandi da usare nel terminal (in una sola volta):

    .\vido-preproc.exe -cmd createindex; webgen

## Nuovo Post
In visual code Terminal:

    cd .\post-src\
    touch 2023-05-30-NuovoPost.txt
    cd ..
Ora edita il nuovo post con Visual code e aggiorna il sito con:

    $env:path = "D:\ruby\ruby_3_2_0\bin;" + $env:path
    .\vido-preproc.exe -cmd createindex; webgen


## Editare l'ultimo post
Si usa il comando (se il browser è già aperto):

    .\vido-preproc.exe --uicmd last --nobrowser

Siccome l'ultimo post è Welcome, non riesce ad editare nient'altro che quello.

## Cambiare altre pagine
Per cambiare le altre pagine, basta usare Visual Code. Tutti i files sono in UTF-8 e _webgen_ li supporta.

Si può usare anche la command line per creare tutti i files index_xx.page con:
    
    vido-preproc.exe -cmd createindex 
    
oppure usando il browser con il quale si è editato il nuovo post.
Ora tocca a webgen (nota che webgen può essere chiamato all'interno del Preprocessor)

## Webgen Versione 1.7.2

    $env:path = "D:\ruby\ruby_3_2_0\bin;" + $env:path
La versione di Ruby 2.3.1 con Windows 11 su _minitoro_ non va. Allora sono passato alla versione 3.2.0,
che non ha funzionato subito. Per trovare gli errori ho usato il debugger su una versione standalone
su D:\scratch\ruby\webgen\stand2 perfettamente funzionante (vedi il suo passaggi.md).
Ho dovuto risolvere il problema con l'Alias e il require di RedCloth. In entrambi i casi ho cambiato
i files dei gems (redcloth.rb e bundle_loader.rb in webgen). Quindi se fai un reinstall del gem, esso non
funzionerà più. In ogni modo bisogna mettere un patch su redcloth e su Webgen.

## Crea un nuovo post col Preprocessor
Si usa vido-preproc.exe. Vedi --usage per il dettaglio. Un nuovo post si aggiunge e si edita
automaticamante con:

    vido-preproc.exe --uicmd new --title "Metti un titolo interssante"
Nel browser (localhost:4200) che parte automaticamente si può editare e salvare il post.
Poi si può creare i files index page usando il link nell'editor.
(Sorgente di vido-preproc.exe messo in D:\scratch\go-lang\vido-preproc)
Per editare l'ultimo post si usa:

    vido-preproc.exe --uicmd last
Vedi tutti i comandi a disposizione con:

    ./vido-preproc.exe --help
Per editare il post corrente senza riaprire una nuova finestra nel browser:

    vido-preproc.exe --uicmd last --nobrowser

## Preprocessor
La teoria è quella di usare il programma 
D:\scratch\go-lang\vido-preproc> go run .\main.go -cmd createindex
questo per generare automaticamente tutti i files inde_xx.page con il loro
navigatore delle pagine settato a modo partendo dai files di post. Ogni post è un file e 
vanno messi assieme in un file di index.
Così basta creare un nuovo file con il post dentro che viene automaticamente 
inserito nelle index.page senza editarle direttamente.
Consiglio di usare 
vido-preproc.exe --uicmd  <comando>
per editare il post e lanciare i vari comandi successivi dal browser, incluso webgen.
Le pagine che avevo editato durante gli anni, le ho splitatte in signoli files usando:
go run .\main.go -cmd splitpages  (NOTA: questo comando non serve più è stato un comando da lanciare una sola volta)
Il risultato si trova sotto: D:\scratch\go-lang\vido-preproc\data\post-src

La configurazione di vido-preproc.exe si trova nel file config.toml di questa directory e i templates
nella sottodirectory templates.
In questa dir l'unica cosa da cambiare è il file config.toml nel caso si voglia usare webgen da una
directory differente oppure l'out è rediretto sotto un'altra dir.

## Update Preprocessor (vido-preproc.exe nuova versione)
Dal progetto vido-preproc lanciare .\start_publish.ps1. Il resto viene fatto da se.

Il programma vido-preproc è un work in progress quindi è possibile che venga aggiornato.
Una volta creato una nuova versione in go nella la dir dove si trova il progetto (golang/vido-preproc),
si crea lo zip con deploy.exe. Questo zip si mette nella sottodirectory ./zips.
Ora si lancia in powershell lo script update-vidopre.ps1 seguendo attentamente i messaggi che compaiono al prompt.
La precendente versione viene memorizzata in ./old

Il progetto vido-preproc.exe si trova su https://github.com/aaaasmile/vido-preproc


## Deploy di invido.it
Con il nuovo server dedicato ho aggiunto https e cambiato il coding, tornando a UTF-8.
Oltre ad aver convertito tutti i files .page in UTF-8 con notepad++, ho cambiato anche 
il file metainfo di webgen per avere UTF-8.
Ora si può editare i files in Visual Code. Per il deploy sul server viene usato il comando linux rsync 
in WLC. 
Per la sincronizzazione dei fiel usa:
./sync_site_invido.sh
Il comando al suo interno è:
rsync -av /mnt/d/Projects/GItHub/InvidositeHtmlgit/out/snapshot/ <user>@invido.it:/var/www/invido.it/html

## Aggiungi un link alla menubar
Occorre un nuovo file con il suffisso .virtual. Poi la referenza va messa nel file metainfo (--- alcn  ).

## WLC deploy 
Usa solo questo comando:
./sync_site_invido.sh

# Sezione Obsoleta

Tengo questa sezione solo per avere dei riferimenti quando si aggiorna un server o i tools dietro alla
creazione del sito. In un futuro potrebbero tornare utili se cambio server o webgen.

### WLC Webgen (obsoleto)
Il resto si può ignorare:
[NOTA, meglio usare webgen sotto windows e WLC solo per il deploy]
Riesco a generare il sito e fare il deploy usando WLC. Questa è la configurazione usata:
igors@Laptop-Toni:/mnt/d/Projects/GItHub/InvidositeHtmlgit$ ruby -v
ruby 2.4.4p296 (2018-03-28 revision 63013) [x86_64-linux]
igors@Laptop-Toni:/mnt/d/Projects/GItHub/InvidositeHtmlgit$ webgen --version
webgen 1.5.1

## Webgen Versione 1.4.1

- con powershel lancia webgen:
$env:path = "D:\ruby\ruby_2_3_1\bin"
o meglio:
$env:path = "D:\ruby\ruby_2_3_1\bin;c:\Program files\git\bin"
webgen
- Ora apri WLC (window linux console) e lancia 
./sync_site_invido.sh

Versione di webgen usata:
PS D:\Projects\GItHub\InvidositeHtmlgit> webgen --version
webgen 1.4.1

## Deploy (non più attuale)

Quello che bisogna fare per aggiornare Invido.it (Sito  http://kickers.fabbricadigitale.it)

vai nella dir D:\Projects\GItHub\InvidositeHtmlgit e apri powershell

    $env:path = "D:\ruby\ruby_2_3_1\bin;c:\Program files\git\bin"

vedi se ruby -v funzione e anche webgen.
Ora con notepad++ modifica i files .page della sottodirectory src
Ricorda di usare la codifica ANSI e mai la UTF8.
Lancia webgen e vedi il risultato nella directory D:\scratch\git-ftp\invidosite
Quando il risultato locale va bene, aggiorna il sito su http://kickers.fabbricadigitale.it/index.html?year=2018
in questo modo:
WLC (con powershell non va) vai nella dir 

    cd /mnt/d/scratch/git-ftp/invidosite/
    git add .
    git commit -m "update"
    git ftp push


In definitiva sono due repository git. La prima che gestisce i sorgenti sotto la dir src ed usa github come repository, 
la seconda è, invece, la sottodir out.
La prima è su github, la seconda è un ftp su http://kickers.fabbricadigitale.it.

## Data 2018

Ho aggiornato webgen alla versione 1.4.1 e funziona. Così come ruby 2.3.1

    $env:path = "D:\ruby\ruby_2_3_1\bin"
Quello che non mi ha funzionato sono gli accenti nel file default.template.
Ricorda che webgen genera la directory di out in D:/scratch/git-ftp/invidosite.
Essa va poi aggiornata con ftp su fabbricadigitale usando WLC (vedi sezione deploy). Non c'è bisogno di copiare nulla.

### Data 30.01.2018
Per usare Visual Studio Code, bisogna fare in modo di impostare l'ambiente di lavoro come Windows 1252 (Einstellungen -> Einstellungen -> "files.encoding":"windows1252"). 
Così files rimangono in ansi e webgen funziona (solo se però il file originale è in iso-8859-1 e non è stato salvato erroneamente in UTF-8).
Con powershell uso:

    $env:path += "D:\ruby\ruby_2_3_0\bin"
Poi ho lanciato:

    webgen
Ho fatto una prova con ruby 2.3.1 dove ho installato webgen 1.4.1 che come sempre 
non ha funzionato. Quindi usa webgen 1.4.0 con ruby 2.3.0

## Deploy (obsoleto)
WLC vai nella dir 

    cd /mnt/d/scratch/git-ftp/invidosite/
    git add .
    git commit -m "update"
    git ftp push

### background git-ftp
Con la WLC su win10 ho provato a fare il deploy usando git-ftp (https://github.com/git-ftp/git-ftp)
Ho crato una nuova repository locale su D:\scratch\git-ftp\invidosite
copiando tutti files della dir /out ed aggiungendoli alla repository (git init invidosite; git add .;git commit -m "update").
Poi ho configurato git-ftp come da help, conl'url, user e password.
Per il deploy ho usato il comando (che serve all'inizio dopo aver settato  url):
git ftp init
Successivamente per ogni cambio e deploy si usa, dopo il commit,
git ftp push
ATTENZIONE: l'url deve essere corretta ed ?unga.
Ho cambiato la destinazione mettendo in web.config:
destination: ["file_system",'D:/scratch/git-ftp/invidosite']
Per vedere tutti i config:
webgen show config

### 03.09.2016
Uso webgen webgen-1.4.0, che ho installato con ruby 2.3.0.
Comando per inizializzare ruby: 
set PATH=D:\ruby\ruby_2_3_0\bin;%PATH%
D:\Projects\GItHub\InvidositeHtmlgit>ruby -v
ruby 2.3.0p0 (2015-12-25 revision 53290) [i386-mingw32]

Poi nella dir si lancia
D:\Projects\GItHub\InvidositeHtmlgit>webgen
Sistema usato: Laptop-Toni


### 13.10.2016
I files del sito invido.it vanno tenuti in ISO-8859-1 per via del server apache di invido.it.
Quindi è meglio evitare di usare Visual Studio Code che converte tutto in UTF8 appena si salva.
Notepad++, invece, va beissimo.

### 09.05.2016
RedCloth sotto windows va installato usando devkit:
gem install RedCloth --platform=ruby
Poi va copiato la lib redcloth_scan.so nella directory 2.3 da creare
D:\ruby_2_3_0\lib\ruby\gems\2.3.0\gems\RedCloth-4.3.0\lib\2.3
Se si usa ruby 2.2.4, il file che si trova in
D:\ruby_2_3_0\lib\ruby\gems\2.3.0\gems\RedCloth-4.3.0\lib\redcloth_scan.so 
va messo nella nuova directory (da creare manualmente):
D:\ruby_2_3_0\lib\ruby\gems\2.3.0\gems\RedCloth-4.3.0\lib\2.2
 
Con ansicon si può anche colorare la cmd sotto windows.
Basta copiare i files released da https://github.com/adoxa/ansicon/releases 
che va e lanciare 
ansicon -i

### 01.05.2016
Ho aggiornato a webgen 1.4.0 e ruby 2.2.4. Il problema principale è il
cambiamento del file webgen.config e del file metainfo nella directory src.
Con 1.4.0 si passa a UTF8 di default, ma tutti i miei sorgenti sono in ISO e 
anche l'output deve essere ISO. Questo si ottiene mettendo il file open iso in
metainfo. Anche la pipeline di redcloth va messa in metainfo.
Il file metainfo è fondamentale, mentre webgen.config serve solo per settare la lingua.
I menu sono cambiati pesantemente e vanno adattati.
Se un file va copiato, ma non viene fatto di default allora va messo in metainfo con l'handler settato a copy
Per i sort va usata la chiave sort_info.
Per avere la directory cuperativa nell'elenco dei progetti nel  menu a sinistra, usare metainfo con 
la chiave in_menu: true.
Per i menu su dispositivi mobile, ho usato lo js slicknav
Su windows per avere i colori nella console, basta copiare i files di 
https://github.com/adoxa/ansicon v1.66 releases in un posto inserito nel path.
Meglio usare sempre il comando 'webgen -v' per avere più informazioni se qualcosa non quadra.


### 6.03.2016
Provato su ubuntu 14.04. Su questo sistema uso rbenv per selzionare la versione di ruby, nel
mio caso la ruby 1.8.7 (2013-12-22 patchlevel 375) [i686-linux].
Bisogna installare la versione 0.5.17. Prima però bisogna installare una versione
di facets compatibile, vale a dire la 2.4.1, e redcloth questo con:

    gem install facets --version 2.4.1
    gem install RedCloth
ora si lancia

    gem install webgen --version 0.5.17

Il programma webgen va poi copiato da ~/.rbenv/versions/1.8.7-p375/bin in 
~/.rbenv/shims

Questi sono gli altri gems che sono stati installati, li metto come referenza
nel caso, come successo con facets, che le nuove versioni non siano più compatibili
con quelle vecchie.

    Fetching: RedCloth-4.2.9.gem (100%)
    Fetching: cmdparse-2.0.6.gem (100%)
    Fetching: maruku-0.7.2.gem (100%)
    Fetching: rake-10.5.0.gem (100%)
    Fetching: rack-1.6.4.gem (100%)
    Fetching: ramaze-2008.06.gem (100%)
    Fetching: kramdown-1.10.0.gem (100%)

L'editor da usare per i files .page è Visual Studio Code. I files .page vanno riaperti in
formato ISO 8859-1. Questo per due motivi, il primo è che webgen 0.5.17 non sa cosa sia UTF8,
il secondo è che il server su fabbrica digitale usa ISO 8859-1.

### 24.01.2016
Quando si inserisce un nuovo file .page, bisogna prestare attenzione che sia in formato ANSI,
altrimenti webgen con ruby 1.8.6 non funziona. Notepad++ e Visual Studio generano file in UTF8 per default
e vanno convertiti in ANSI esplicitamente.

### 15.12.2014
Ho aggiornato il sito usando la repository su github, ho usato webgen 0.5.17
sotto windows 7. Con la codifica utf8 non riesco a fare andare i files html su fabbrica digitale.
Tutti i caratteri accentati sono resi male. Allora con notepad++ ho cambiato la codifica
di tutti i sorgenti *.page usando il formato ANSI (Menu Encoding -> Convert to Ansi).
Il problema c'è anche nel manuale della Cuperativa. Quindi quando verrà rigenerato, per
piacere cambia il formato da UTF a iso-8859-1.

### 22.11.2014
Sto sempre usando webgen 0.5.17 sotto windows e ho messo redcloth in config.yaml
Ho tolto tutti i riferimenti a progetti vecchi come briscola_net e ho spostato
tutto su kickers.fabbricadigitale.it


### 10.12.2012
Sto usando webgen 0.5.17 che ho installato manualmente scaricando tutti
i pacchetti gem necessari su ruby 1.8.6.

### Dal 01.2012
Ho convertito con molta difficoltà il progetto per webgen 0.5.14.
Se con Notepad++ crei una page, bisogna stare attenti che il formato
sia quello giusto. Bisogna settare su: encode in UTF 8 without BOM.
Non ancora chiaro del tutto come usare il sort_info per i menu che sono
in sotto directory, così com'è ora funziona solo per i files che si trovano
nella stessa directory, quelli della sottodirectory vengono messi in fondo.

### Info generali

Per generare la home page dell'invido usare in questa directory:
webgen

La home page all'inizio è stata creata col comando
webgen create invido

Poi nella directory invido ho cambiato lo stile su andreas09
webgen use website_style andreas09

Ho cambiato il file default.css in mod da avere i submenu formattati come
la home page di webgen. 

Nella directory ho creato config.yaml e metainfo.yaml in modo da cambiare delle
funzionalità di default. metainfo.yaml è usato per cambiare l'ordine dei menu.
config.yaml è usato per cambiare la lingua di default.

Per quanto riguarda il formato da usare esso è redcloth, textile. Il riferimento
si trova su :
http://hobix.com/textile/

L'esempio della home page di webgen, vale a dire il sorgente della home page
webgen, si trova su:
/usr/local/lib/ruby/gems/1.8/gems/webgen-0.4.7/doc/src
L'esempio è molto istruttivo e serve per capire come i vari tags funzionano.
La home è su http://webgen.rubyforge.org

### Problemi incontrati
* Nomi dei files relativi non funzionano. Sol: usare relocatable. Per esempio
{relocatable: /images/invido_1.gif}

* Creare una lista bullet. Bisogna avere tutti gli asterischi della lista sulla
stessa colonna. Se si vuole una lista delle liste si lascia una riga vuota e si mette
l'asterisco all'interno usando sempre la stessa colonna per l'asterisco

### Server su fabbrica digitale non supporta utf8
Esso supporta solo ISO-8859-1. Se i files .template e .page sono in formato utf8,
allora anche il risultato sarà in utf8. Se però lo si legge con ISO-8859-1 allora
si leggerà tutto male. Bisogna avere allora tutti i files sorgenti come ISO-8859-1.
Per conventire if files da utf8 a iso-8859-1 ho usato komodo su ogni file
Si potrebbe usare, parlo solo a livello teorico, il seguente codice:
require 'iconv'
content_latin1 = Iconv.new('iso-8859-1', 'utf-8').iconv(file_content_utf8)

### Aggiornamento 22.11.2014 sul problema utf8
Sotto windows web gen non riesce a generare i files usando iso-8859-1, li vuole in 
UTF 8 without BOM. Allora non mi è rimasto altro che convertire i files in out in iso-8859-1
che notepad++ si converte prima in utf8


### Server su Hosteurope
NOTA: L'Ip di invido.it è al 2012: 46.163.74.64

Se si vuole avere questi files html sul server hosteurope, allora bisogna
avere i files in utf8. Come ho fatto a farlo andare? Basta caricare i
files di invido.it su 92.51.146.97 usando ftp con l'utente del dominio
di default, nel caso della prova igorvivace col dominio vivacesoft.com.
Da notare che lasciando i files di default che ha creato Plesk, usando
http://92.51.146.97/ non si vede nulla. Cancellando questi files e 
caricando i miei invece ha funzionato. La prima volta ho anche dovuto
eseguire un restart di apache.
Ora però il problema al contrario, qui serve utf8. Allora ho rimesso
tutto in utf8, templates e charset ed encodind usando file preferences
in komodo.
Poi per aggiornare i files su hosteurope basta usare il seguente comando:

rsync -raz --progress --size-only /home/igor/Projects/invidosite/invido08webgen/out/* /var/www/vhosts/lvps46-163-74-64.dedicated.hosteurope.de/site1_invido

Con il nuovo server il link dei documenti html è:
/var/www/vhosts/lvps46-163-74-64.dedicated.hosteurope.de/site1_invido
Mentre l'indirizzo IP è: 46.163.74.64


### Update sul sito direttamente
Usando il sito http://shiftedit.net/home si possono cambiare i files del sito
poi usando la console si lancia webgen nella directory
/home/igor/Projects/invidosite/invido08webgen
Poi si sincronizza il contenuto dell'output con quello per il sito usando il comando:
sudo ./sync_site_invido.sh 
ATTENZIONE: in questo modo si aggiorna il sito senza vedere i cambiamenti in modo locale

### Auguri di Natale 2012
In questa directory:
~/Projects/node/nodeflakes/server_install/server
lanciare ./start_all.sh
Poi modificare il file default.template e scommentare la parte che si collega a nodeflakes (css e javascript)
Ora lanciare webgen a linea di comando e poi sincronizzare i files
sudo ./sync_site_invido.sh 


