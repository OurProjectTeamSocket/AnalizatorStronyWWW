# AnalizatorStronyWWW
Zlecenie dał typa z Fiverr'a w Go

<h2> Czym jest program </h2>

<h4> Program ma byc stroną www na której można się zarejestrować i sprawdzać status swojego serwera </h4>

<h2> Co ma robic program? </h2>
<h4> <strike> Program ma odbierać połączenia przez klientów, monitorowac ich zachowania na stronie i robić z nich wykresy </strike> </h4>
<h4> Edit: Program co 5 minut ( w bezpłatnej wersji ) lub co 1 minute ( w płatnej wersji ) ma wysyłać zapytanie do strony WWW którą user poda i sprawdzać czy działa czy nie i w jakim czasie udziela odpowiedzi. Komunikat o działaniu ma wyświetlac pod spodem a odpowiedzi ma wyświetlac w postaci diagramu </h4>
</br>
<h2> Co będzie nam do tego potrzebne </h2>
<h4> <li> Go </li>
     <li> GORM </li>
     <li> <strike> Strona testowa </strike> Front end </li>
     <li> Serwer HTTP </li> </h4>
     
<h2> Program w punktach </h2>
<h4> <li> Moduł rejestracji </li>
     <li> Moduł logowania ( tez przez google mail'a ) </li>
     <li> Moduł przywracania hasła </li>
     <li> Strona głowna ( i jedyna ) </li>
     <ul>
          <li> Pasek po lewej: zawierający podpięte strony ( buttony ) i zdjęcie user'a, jego nazwę i status </li>
          <li> Pasek na górze: który ma w sobie dwie informacje Host URL ( string ) i SSL ( bool ) </li>
          <li> Główny kontener: wykresy i informacje o stronie - status strony, of kiedy strona stoi, czy ma SSL i wykresy. Kolejno: Monitor działania strony, monitor odpowiedzi, zastawienie procentowe aktywności strony. Na samym dole sa notyfikacja - nie wiem o co z nimi chodzi. </li>
     </ul>
     <li> Najebanie jak najwięcej grafów jak się da </li> </h4>
     
<h2> Materiały wysłane przez klienta <h2>
<h4> Lista próśb:
     <li> 1. Use uuid instead of id incremental postgresql. </li>
     <li> 2. Use dotenv to keep secrets. </li>
     <li> 3. Kindly have thin controller actions. </li>
     <li> 4. Background jobs plugins has a prebuilt dashboard then use that gopackage. </li>
     <li> 5. Make environment configurable like development, staging, production. </li>
     <li> 6. I will create a repo and will invite you once you pushed a basic repo please inform me I will do CI/CD deployment. </li>
     <li> 7. If you can built any kind of graphs from similar sites please include it. </li>
     <li> 8. Kindly check wireframe shared in detail. </li>
     <li> 9. Try some similar site Analytics graphs. </li>
     <li> 10. If possible page HTML also store somewhere so in future we can give them seo Analytics also. </li>
     <li> 11. Whatever the values you can grab grab it on each request. </li>
     <li> 12. Save DB space maximum as possible. </li>
     Rysunki: </h4>
     ![jeden](https://user-images.githubusercontent.com/31569763/128310697-19918558-9542-4931-9461-6d5c240b0130.jpeg)
     ![dwa](https://user-images.githubusercontent.com/31569763/128310709-8e3fe977-5b23-4390-aa8d-5441c9765b0e.jpeg)


