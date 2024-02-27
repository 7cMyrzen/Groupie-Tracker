//----------------------------//
// Pour chaque touche appuyée //
//----------------------------//

function checkKeyPress(event) {
  HideAndShow();

  var keyCode = event.keyCode || event.which;

  if (keyCode === 13) {
    search();
  }
}

//-------------------------------//
// Chercher la page de l'artiste //
//-------------------------------//

function search() {
  // Récupérer la valeur de la barre de recherche
  var searchValue = document.getElementById("search").value.toLowerCase();

  // Parcourir tous les éléments avec la classe 'card-container'
  var cardContainers = document.querySelectorAll(".card-container");

  var groupFound = false;

  cardContainers.forEach(function (card) {
    // Récupérer l'ID du groupe
    var groupId = card.id;

    // Récupérer le nom du groupe
    var groupName = card.querySelector(".cardname").textContent.toLowerCase();

    // Vérifier si le groupe recherché existe dans le nom du groupe
    if (groupName.includes(searchValue)) {
      console.log("Nom du groupe: " + searchValue + ", ID: " + groupId);
      // Rediriger vers la nouvelle URL en utilisant window.location.href
      window.location.href = "http://localhost:8080/groupe/" + groupId;
      groupFound = true;
    }
  });

  // Afficher un message si le groupe n'a pas été trouvé
  if (!groupFound) {
    console.log("Le groupe '" + searchValue + "' n'a pas été trouvé.");
  }
}

//---------------------------------//
// Cacher et afficher les éléments //
//---------------------------------//

function HideAndShow() {
  // Récupérer la valeur de la barre de recherche
  var searchValue = document.getElementById("search").value.toLowerCase();

  //Cacher les éléments qui ne correspondent pas à la recherche et afficher ceux qui correspondent
  var cardContainers = document.querySelectorAll(".card-container");
  cardContainers.forEach(function (card) {
    // Récupérer le nom du groupe
    var groupName = card.querySelector(".cardname").textContent.toLowerCase();
    if (groupName.includes(searchValue)) {
      card.style.display = "block";
    } else {
      card.style.display = "none";
    }
  });
}
