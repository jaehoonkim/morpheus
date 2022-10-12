INSERT INTO `template` (`uuid`, `name`, `summary`, `origin`, `created`) VALUES ('00000000000000000000000000000036', 'kubernetes_services_patch', 'class=\'kubernetes\' resource=\'services\' verb=\'patch\' group_version=\'v1\'', 'predefined', NOW());
INSERT INTO `template_command` (`uuid`, `name`, `summary`, `template_uuid`, `sequence`, `method`, `args`, `result_filter`, `created`) VALUES ('00000000000000000000000000000036', 'kubernetes_services_patch_0', 'class=\'kubernetes\' resource=\'services\' verb=\'patch\' group_version=\'v1\'', '00000000000000000000000000000036', '0', 'kubernetes.services.patch.v1', '{"type":"object","properties":{"namespace":{"type":"string","pattern":"^."},"name":{"type":"string","pattern":"^."},"patch_type":{"type":"string","enum":["json","merge"]}},"required":["namespace","name","patch_type"],"dependentRequired":{"patch_type":["patch_data"]},"additionalProperties":false,"if":{"properties":{"patch_type":{"const":"json"}}},"then":{"properties":{"patch_data":{"type":"array","items":{"oneOf":[{"additionalProperties":false,"required":["value","op","path"],"properties":{"path":{"$ref":"#/$defs/path"},"op":{"description":"The operation to perform.","type":"string","enum":["add","replace","test"]},"value":{"description":"The value to add, replace or test."}}},{"additionalProperties":false,"required":["op","path"],"properties":{"path":{"$ref":"#/$defs/path"},"op":{"description":"The operation to perform.","type":"string","enum":["remove"]}}},{"additionalProperties":false,"required":["from","op","path"],"properties":{"path":{"$ref":"#/$defs/path"},"op":{"description":"The operation to perform.","type":"string","enum":["move","copy"]},"from":{"$ref":"#/$defs/path","description":"A JSON Pointer path pointing to the location to move/copy from."}}}]}}}},"else":{"properties":{"patch_data":{"type":"object"}}},"$defs":{"path":{"description":"A JSON Pointer path.","type":"string"}}}', NULL, NOW());

INSERT INTO `template_recipe` ( `method`, `args`, `name`, `summary`, `created`) VALUES ('kubernetes.services.patch.v1', '{"type":"object","properties":{"namespace":{"type":"string","pattern":"^."},"name":{"type":"string","pattern":"^."},"patch_type":{"type":"string","enum":["json","merge"]}},"required":["namespace","name","patch_type"],"dependentRequired":{"patch_type":["patch_data"]},"additionalProperties":false,"if":{"properties":{"patch_type":{"const":"json"}}},"then":{"properties":{"patch_data":{"type":"array","items":{"oneOf":[{"additionalProperties":false,"required":["value","op","path"],"properties":{"path":{"$ref":"#/$defs/path"},"op":{"description":"The operation to perform.","type":"string","enum":["add","replace","test"]},"value":{"description":"The value to add, replace or test."}}},{"additionalProperties":false,"required":["op","path"],"properties":{"path":{"$ref":"#/$defs/path"},"op":{"description":"The operation to perform.","type":"string","enum":["remove"]}}},{"additionalProperties":false,"required":["from","op","path"],"properties":{"path":{"$ref":"#/$defs/path"},"op":{"description":"The operation to perform.","type":"string","enum":["move","copy"]},"from":{"$ref":"#/$defs/path","description":"A JSON Pointer path pointing to the location to move/copy from."}}}]}}}},"else":{"properties":{"patch_data":{"type":"object"}}},"$defs":{"path":{"description":"A JSON Pointer path.","type":"string"}}}', 'kubernetes-services-patch-v1', 'namespace:required;string;,name:required;string;,patch_type:required;string;,patch_data:required;string;', NOW());
