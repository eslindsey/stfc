package stfc

import (
	"bytes"
	"encoding/json"
)

type FleetModifyFleetRequest struct {
	FleetId    uint64   `json:"fleet_id"`
	ShipLayout []uint64 `json:"ship_layout"`
}

func (s *Session) FleetModifyFleet(request *FleetModifyFleetRequest) ([]byte, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	_, err = s.Post("/fleet/modify_fleet", nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	// TODO: decode proto & JSON, return
	return nil, nil
/*
[message]    1                                                                                                                 
[uint32]     1.1         42                                                                                                    
[string]     1.2         {"my_deployed_fleets": {}, "static_update": null, "fleets": {"1936358958516484622": {"ship_ids": [1659671079071073774], "name": "", "drydock_id": 18, "last_recall": "1970-01-01T00:00:00", "officers": [3579242146, 909354959, 4150311506, 1689277174, 2702893021, null, null, null, null, null], "stats": {"55": 1.0, "79": 1.2}, "repair_time": 0, "repair_cost": {}, "precalculated_repair": true}, "1936358958533261841": {"ship_ids": [1249481011289736408], "name": "", "drydock_id": 21, "last_recall": "1970-01-01T00:00:00", "officers": [2334517642, 370544309, 1395454844, 2446409392, 3304441016, 1859906553, 1131760724, 180486983, null, null], "stats": {"55": 1.0, "79": 1.0}, "repair_time": 0, "repair_cost": {}, "precalculated_repair": true}, "1936358958499707405": {"ship_ids": [1250223476573553961], "name": "", "drydock_id": 20, "last_recall": "1970-01-01T00:00:00", "officers": [2030703558, 440622131, 1580157615, 743931698, 56060748, 3999633781, 3769449499, 1800794103, null, null], "stats": {"55": 1.0, "79": 1.0}, "repair_time": 0, "repair_cost": {}, "precalculated_repair": true}, "1936358958491318796": {"ship_ids": [1732813461068793054], "name": "", "drydock_id": 17, "last_recall": "1970-01-01T00:00:00", "officers": [3583932904, null, null, null, null, null, null, null, 2959514562], "stats": {"55": 1.0, "79": 1.0}, "repair_time": 0, "repair_cost": {}, "precalculated_repair": true}, "1936358958524873231": {"ship_ids": [1655581509039768818], "name": "", "drydock_id": 19, "last_recall": "1970-01-01T00:00:00", "officers": [329940464, 3990993357, 2765885322, 988947581, 656972203, 2822661458, 1730335425, 2695272429], "stats": {"55": 1.0, "79": 1.0}, "repair_time": 0, "repair_cost": {}, "precalculated_repair": true}}}
[message]    2                                                                                                                 
[uint32]     2.1         56                                                                                                    
[message]    2.2                                                                                                               
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     147c73db70414892bf579a99811f0f7f                                                                      
[uint32]     2.2.1.3     86400                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778146                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906336422447946                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     1385c4adb40e45b2977e35fd8c351049                                                                      
[uint32]     2.2.1.3     259200                                                                                                
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778166                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906559567809076                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     ced0bd53dbff4296b6d5a6611d8ec794                                                                      
[uint32]     2.2.1.3     43200                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778209                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906509622037493                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     4                                                                                                     
[string]     2.2.1.2     ac2bc741269b475daa3b0ccd7cde8706                                                                      
[uint32]     2.2.1.3     145180                                                                                                
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1686685978                                                                                            
[uint32]     2.2.1.5     11194                                                                                                 
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.10                                                                                                          
[uint32]     2.2.1.10.1  13                                                                                                    
[uint32]     2.2.1.10.2  30                                                                                                    
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     f506c47439d949b9a7d309fbdd6c7209                                                                      
[uint32]     2.2.1.3     86400                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778260                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906403917187978                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     8b42e25a83ad42999999096d3607574e                                                                      
[uint32]     2.2.1.3     86400                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778286                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906270613818095                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     3d4d52b2402c43d094be9106f78250bd                                                                      
[uint32]     2.2.1.3     172800                                                                                                
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778310                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906607282211440                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     6a10b19816184b09adb1b978adfa6938                                                                      
[uint32]     2.2.1.3     86400                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778337                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941903666135583425                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     09a23517f74146e99bb21bf63f03aee0                                                                      
[uint32]     2.2.1.3     86400                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778357                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941903666479516355                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     3                                                                                                     
[string]     2.2.1.2     29bbfa2ff8aa4131bc1e6563d5d0ea4f                                                                      
[uint32]     2.2.1.3     53753                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1686685609                                                                                            
[uint32]     2.2.1.5     4124                                                                                                  
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.15                                                                                                          
[uint32]     2.2.1.15.1  871758445                                                                                             
[uint32]     2.2.1.15.2  1                                                                                                     
[message]    2.2.1.18                                                                                                          
[uint64]     2.2.1.18.1  18446744073709551615                                                                                  
[string]     2.2.1.18.2  x5e60073fe45433ba86e95a8b80138af                                                                      
[uint32]     2.2.1.18.3  859804212                                                                                             
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     f209b54955a347b48684af1e0efed01f                                                                      
[uint32]     2.2.1.3     172800                                                                                                
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778375                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941903666135583426                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     13                                                                                                    
[string]     2.2.1.2     e03288a8eb6d4e93bba2f2c2e3147119                                                                      
[uint32]     2.2.1.3     43200                                                                                                 
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1685778399                                                                                            
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.17                                                                                                          
[uint64]     2.2.1.17.1  1941906456698309564                                                                                   
[message]    2.2.1                                                                                                             
[uint32]     2.2.1.1     4                                                                                                     
[string]     2.2.1.2     466c858cb61d44deb0c3ce444193b16c                                                                      
[uint32]     2.2.1.3     573639                                                                                                
[message]    2.2.1.4                                                                                                           
[uint32]     2.2.1.4.1   1686685995                                                                                            
[uint32]     2.2.1.5     44295                                                                                                 
[fixed32]    2.2.1.6     1065353216                                                                                            
[string]     2.2.1.7     x5e60073fe45433ba86e95a8b80138af                                                                      
[message]    2.2.1.10                                                                                                          
[uint32]     2.2.1.10.1  40                                                                                                    
[uint32]     2.2.1.10.2  34                                                                                                    
[message]    2                                                                                                                 
[uint32]     2.1         79                                                                                                    
[message]    3                                                                                                                 
[uint32]     3.1         1686686683                                                                                            
[string]     4           x5e60073fe45433ba86e95a8b80138af                                                                      
[string]     5           v12.3.0-M55.000000115-universe_package_708d238e_M55.0.0-RC6_fixture_M55.000000115                     
*/
}


