

$(document).ready(function() {
  var filters = []
  filters['meal'] = []
  filters['hall'] = []
  filters['filter'] = []

  // add or remove filter from list
  function toggleFilter(type, value) {
    console.log(type)
    console.log(filters[type])
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
    if ($(id).hasClass('highlight'))
      $(id).removeClass('highlight');
    else
      $(id).addClass('highlight');
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

  // initial setup
  getCurrentMeal();
  renderList();

  // upvote clicked
  $('.upvote').click(function() {
    var id = $(this).data('food-id')
    var vote = $('#' + id).find('.votes')
    vote.html(parseInt(vote.text())+1);
    $.post( "/api/inc/" + id);
  });

  // downvote clicked
  $('.downvote').click(function() {
    var id = $(this).data('food-id')
    var vote = $('#' + id).find('.votes')
    vote.html(parseInt(vote.text())-1);
    $.post("/api/dec/" + id);
  });

  // filter selected
  $('.circle').click(function() {

    var type = $(this).attr('class');
    var value = $(this).attr('id')

    type = $.trim(type.replace('circle', '').replace('highlight', ''));

    console.log(filters[type]);

    highlightFilter(value)
    toggleFilter(type, value);

    renderList();   
  });

  $(function() {
    var id2convert = "meals",
        divs = $("div", "#" + id2convert),
        select = $("<select id=\"" + id2convert + "\">");
    
    divs.each(function() {
        var text = $(this).text();
        select.append("<option value=\"" + text + "\">" + text + "</option>");
    });
    
    divs.parent().replaceWith(select);
  });
  $(function() {
    var id2convert = "halls",
        divs = $("div", "#" + id2convert),
        select = $("<select id=\"" + id2convert + "\">");
    
    divs.each(function() {
        var text = $(this).text();
        select.append("<option value=\"" + text + "\">" + text + "</option>");
    });
    
    divs.parent().replaceWith(select);
  });
  $(function() {
    var id2convert = "diets",
        divs = $("div", "#" + id2convert),
        select = $("<select id=\"" + id2convert + "\">");
    
    divs.each(function() {
        var text = $(this).text();
        select.append("<option value=\"" + text + "\">" + text + "</option>");
    });
    
    divs.parent().replaceWith(select);
  });
});


(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-62469249-1', 'auto');
ga('send', 'pageview');

