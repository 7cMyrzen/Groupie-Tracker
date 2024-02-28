//---//
//---//
//---//

function updateSelection1(selectedOption) {
  if (selectedOption === "option1") {
    // Si l'option 1 est sélectionnée, décocher l'option 2
    document.getElementById("option1-2").checked = false;
  } else if (
    selectedOption === "option2" &&
    document.getElementById("option1-1").checked
  ) {
    // Si l'option 2 est sélectionnée et l'option 1 était déjà sélectionnée, décocher l'option 1
    document.getElementById("option1-1").checked = false;
  }
}

function updateSelection2(selectedOption) {
  if (selectedOption === "option1") {
    // Si l'option 1 est sélectionnée, décocher l'option 2
    document.getElementById("option2-2").checked = false;
  } else if (
    selectedOption === "option2" &&
    document.getElementById("option2-1").checked
  ) {
    // Si l'option 2 est sélectionnée et l'option 1 était déjà sélectionnée, décocher l'option 1
    document.getElementById("option2-1").checked = false;
  }
}

function updateSelection3(selectedOption) {
  if (selectedOption === "option1") {
    // Si l'option 1 est sélectionnée, décocher l'option 2
    document.getElementById("option3-2").checked = false;
  } else if (
    selectedOption === "option2" &&
    document.getElementById("option3-1").checked
  ) {
    // Si l'option 2 est sélectionnée et l'option 1 était déjà sélectionnée, décocher l'option 1
    document.getElementById("option3-1").checked = false;
  }
}

function getSelectedOptions() {
  var o11 = document.getElementById("option1-1").checked;
  var o12 = document.getElementById("option1-2").checked;
  var o21 = document.getElementById("option2-1").checked;
  var o22 = document.getElementById("option2-2").checked;
  var o31 = document.getElementById("option3-1").checked;
  var o32 = document.getElementById("option3-2").checked;

  var selectedOptions = [];

  if (o11) {
    selectedOptions[0] = "option1";
  } else {
    if (o12) {
      selectedOptions[0] = "option2";
    } else {
      selectedOptions[0] = "none";
    }
  }

  if (o21) {
    selectedOptions[1] = "option1";
  } else {
    if (o22) {
      selectedOptions[1] = "option2";
    } else {
      selectedOptions[1] = "none";
    }
  }

  if (o31) {
    selectedOptions[2] = "option1";
  } else {
    if (o32) {
      selectedOptions[2] = "option2";
    } else {
      selectedOptions[2] = "none";
    }
  }

  console.log("Selected options:", selectedOptions);

  return selectedOptions;
}

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

async function filter() {
  var selectedOptions = getSelectedOptions();

  // Select all elements with class "card-container"
  const cardContainers = document.querySelectorAll(".card-container");

  cardContainers.forEach((container) => {
    container.style.display = "block";
  });

  // Loop through each card container
  cardContainers.forEach((container) => {
    // Extract the values from the cardinfo-content elements
    var membersNumber = parseInt(
      container.querySelector(".cardinfo-content1").textContent,
      10
    );
    // make sure to convert the string to a number
    membersNumber = membersNumber * 1;

    const firstAlbumFull =
      container.querySelector(".cardinfo-content2").textContent;
    const firstAlbumLast4Digits = parseInt(firstAlbumFull.slice(-4), 10); // Keep only the last 4 digits
    var creationDate = parseInt(
      container.querySelector(".cardinfo-content3").textContent,
      10
    );
    // make sure to convert the string to a number
    creationDate = creationDate * 1;

    // Update the content with the converted values (optional)
    container.querySelector(".cardinfo-content1").textContent = membersNumber;
    container.querySelector(".cardinfo-content2").textContent =
      firstAlbumLast4Digits;
    container.querySelector(".cardinfo-content3").textContent = creationDate;

    // Output the converted values (optional)
    console.log("Members Number:", membersNumber);
    console.log("First Album (Last 4 digits):", firstAlbumLast4Digits);
    console.log("Creation Date:", creationDate);
    var filter1 = true;
    var filter2 = true;
    var filter3 = true;

    // Check the third filter
    if (selectedOptions[2] === "option1") {
      if (membersNumber > 3) {
        filter3 = false;
      }
    } else if (selectedOptions[2] === "option2") {
      if (membersNumber <= 3) {
        filter3 = false;
      }
    }

    // Check the second filter
    if (selectedOptions[1] === "option1") {
      if (firstAlbumLast4Digits > 2000) {
        filter2 = false;
      }
    } else if (selectedOptions[1] === "option2") {
      if (firstAlbumLast4Digits <= 2000) {
        filter2 = false;
      }
    }

    // Check the first filter
    if (selectedOptions[0] === "option1") {
      if (creationDate > 2000) {
        filter1 = false;
      }
    } else if (selectedOptions[0] === "option2") {
      if (creationDate <= 2000) {
        filter1 = false;
      }
    }

    console.log(
      creationDate,
      firstAlbumLast4Digits,
      membersNumber,
      "/////",
      filter1,
      filter2,
      filter3
    );
    if (filter1 && filter2 && filter3) {
      container.style.display = "block";
    } else {
      container.style.display = "none";
    }
  });
}
