# nutzamqp


	  declare:
	    exchanges:
	      - name: x1
	        type: direct
	        durable: true
	        autodeleted: false
	        
	      - name: x2
	        type: fanout
	        durable: false
	        autodeleted: false
	        
	    queues:
	      - name: q1
	        durable: true
	        delete_when_unused: false
	        
	      - name: q2
	        durable: true
	        delete_when_unused: false
	        
	      - name: q3
	        durable: true
	        delete_when_unused: false
	        
	      - name: q4
	        durable: false
	        delete_when_unused: true
	        
	    bindings:
	      # [exchange, key, queue]
	      - [x1, kq1, q1]
	      
	    xbindings:
	      # [exchange, key, exchange]
	      - [x2, "", x3]
	      
