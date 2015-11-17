shell = require 'shell'

search = $('input.search')
hits   = $('.hits')
result = $('ul.result')

search.on 'keyup', ->
  text = ""
  keyword = $(@).val()
  if keyword != ""
    $.ajax
      url: 'http://localhost:8000/search/' + keyword
      dataType: 'json'
      success: (data) ->
        text = ""
        for path in data['hits']
          text += "<li><div class=\"score-wrap\"><div class=\"score\" style=\"width: #{path['score'] * 75}px;\"></div></div><a data-path=\"#{path['id']}\">#{path['id']}</a></li>"
        if text.length > 0
          total_hits = data['total_hits']
          took = data['took']
          hits.html "#{total_hits}"
          result.html text
          $('a[data-path]').on 'click', ->
            console.log 'aa?'
            shell.showItemInFolder $(@).data('path')
        else
          result.html "No Result"
          hits.html ""
  else
    result.html "Please input keyword"
    hits.html ""
