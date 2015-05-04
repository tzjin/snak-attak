

$(document).ready(function() {
  var filters = []
  filters['meal'] = []
  filters['hall'] = []
  filters['filter'] = []

  var upvotes = []
  var downvotes = []


/********** FILTER MAGIC **********/

  // add or remove filter from list
  function toggleFilter(type, value) {
    var index = filters[type].indexOf(value);
    if (index < 0)
      filters[type].push(value);
    else 
      filters[type].splice(index, 1);
  }

  // find current meal
  function getCurrentMeal() {
    var time = new Date();
    var hour = time.getHours();
    if (hour < 10 )
      filters['meal'].push('b')
    else if (hour < 14)
      filters['meal'].push('l')
    else // if (hour < 20)
      filters['meal'].push('d')

    highlightFilter(filters['meal'][0])
  }

  // highlight selected filter
  function highlightFilter(value) {
    var id = "div[id='" + value + "']"
    $(id).toggleClass('highlight')
  }

  // hide or show each food item in list
  function renderList() {
    $('.food').each(function() {
      var valid = true;
      // check each type of filter
      for (var ls in filters) {
        
        // iterate through each filter value
        var valid_filt = (filters[ls].length == 0);
        for (var idx in filters[ls]) {
          
          // custom logic for pork/nut-free
          if (filters[ls][idx] == 'Pork' || filters[ls][idx] == 'Nuts')
            valid_filt |= !$(this).hasClass(filters[ls][idx]);
          
          // else just check to see if contains filter
          else {
            valid_filt |= $(this).hasClass(filters[ls][idx]);

            // all vegan items are also vegetarian
            if (filters[ls][idx] == 'Vegetarian')
              valid_filt |= $(this).hasClass('Vegan');
          }
        }

        // and logic between filter types
        valid &= valid_filt;
      }

      // hide or show
      if (valid)
        $(this).removeClass('hidden');
      else
        $(this).addClass('hidden');
    });
  }


/********** VOTING MAGIC **********/
  
  // update view and send post request
  function upvote(id) {
    var vote = $('#' + id).find('.votes')
    upvotes.push(id);
    vote.html(parseInt(vote.text())+1);
    $.post( "/api/inc/" + id);
  }

  // update view and send post request
  function downvote(id) {
    var vote = $('#' + id).find('.votes')
    downvotes.push(id)
    vote.html(parseInt(vote.text())-1);
    $.post("/api/dec/" + id);
  }

  // highlight arrow
  function toggleUpvote(id) {
    var vote = $('.upvote[data-food-id="' + id + '"]')
    vote.toggleClass('grey')
  }

  function toggleDownvote(id) {
    var vote = $('.downvote[data-food-id="' + id + '"]')
    vote.toggleClass('grey')
  }


/********** CLICK HANDLERS **********/

  // initial setup
  getCurrentMeal();
  renderList();

  // upvote clicked
  $('.upvote').click(function() {
    var id = $(this).data('food-id')
    
    if (upvotes.indexOf(id) != -1) {
      console.log("Already voted on this item!")
      return
    }

    if (downvotes.indexOf(id) != -1) {
      downvotes.splice(downvotes.indexOf(id), 1)
      toggleDownvote(id)
      upvote(id)
    }

    toggleUpvote(id)
    upvote(id)
  });

  // downvote clicked
  $('.downvote').click(function() {
    var id = $(this).data('food-id')
    
    if (downvotes.indexOf(id) != -1) {
      console.log("Already voted on this item!")
      return
    }

    if (upvotes.indexOf(id) != -1) {
      upvotes.splice(upvotes.indexOf(id), 1)
      toggleUpvote(id)
      downvote(id)
    }

    toggleDownvote(id)
    downvote(id)
  });

  // filter selected
  $('.circle').click(function() {

    var type = $(this).attr('class');
    var value = $(this).attr('id')

    type = $.trim(type.replace('circle', '')
      .replace('highlight', ''));

    console.log(filters[type]);

    highlightFilter(value)
    toggleFilter(type, value);

    renderList();   
  });

  /*$(function() {
    var id2convert = "meals",
        divs = $("div", "#" + id2convert).find('.filter_name'),
        select = $("<select id=\"" + id2convert + "\">");
    divs.each(function() {
        var text = $(this).text();
        select.append("<option value=\"" + text + "\">" + text + "</option>");
    });
    
    divs.parent().parent().replaceWith(select);
  });*/
});


(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-62469249-1', 'auto');
ga('send', 'pageview');

