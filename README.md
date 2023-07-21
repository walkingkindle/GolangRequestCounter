# GolangRequestCounter


Zadatak 3:

Napisati API upotrebom standardne biblioteke ili u frameworku po izboru koji ima sledece mogucnosti:
    * svaki request se broji
    * implementiraj po zelji:
        * uz pomoc globalne varijable
        * mehanizam iz frameworka
    * GET /count ruta koja daje trenutni broj requestova, i ona se broji

Pitanja na koja treba dati odgovore:
    * Na sta treba paziti kada se radi sa globalnim varijablama u konkurentnom okruzenju?

Dodatno (nije obavezno):
    * Go unit testovi
    * Dockerizovati mikroservis

------------------------------------------------------------------------------------------------------------------

    * Na sta treba paziti kada se radi sa globalnim varijablama u konkurentnom okruzenju?
        Uglavnom, da se varijable ne bi slucajno promenile, ili da ne bih doslo do grekse, obicno je moja praksa izbegavati
        globalne varijable. Ako to projekat nalaze, predlazem koriscene nekog package-a, u mojem slucaju mutual exclusion, koji
        ce da se pobrine za to da samo jedna goroutina koristi odredjenu varijablu.
