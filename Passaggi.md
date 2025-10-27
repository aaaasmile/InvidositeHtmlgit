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


