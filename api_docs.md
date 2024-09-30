
# Documentação da API do VetBlock

**Base URL:** `/api/v1`

## Endpoints disponíveis:

### 1. Adicionar Animal
- **Rota:** `POST /animals`
- **Descrição:** Adiciona um novo animal no sistema.

#### Corpo da Requisição:
```json
{
  "name": "Rex",
  "species": "Canine",
  "breed": "Golden Retriever",
  "weight": 30.5,
  "age": 3,
  "description": "Animal dócil",
  "cpf_tutor": "12345678901"
}
```

#### Resposta de Sucesso:
- **Código:** 201 Created
- **Corpo:**
```json
{
  "id": "UUID gerado",
  "name": "Rex",
  "species": "Canine",
  "breed": "Golden Retriever",
  "weight": 30.5,
  "age": 3,
  "description": "Animal dócil",
  "cpf_tutor": "12345678901"
}
```

#### Possíveis Erros:
- 400 Bad Request: Corpo da requisição inválido.
- 500 Internal Server Error: Falha ao salvar o animal.

---

### 2. Atualizar Animal
- **Rota:** `PUT /animals/:id`
- **Descrição:** Atualiza os dados de um animal existente.

#### Parâmetros de URL:
- `id`: UUID do animal.

#### Corpo da Requisição:
```json
{
  "name": "Rex",
  "species": "Canine",
  "breed": "Golden Retriever",
  "weight": 35,
  "age": 4,
  "description": "Animal muito ativo",
  "cpf_tutor": "12345678901"
}
```

#### Resposta de Sucesso:
- **Código:** 200 OK
- **Corpo:**
```json
{
  "id": "UUID do animal",
  "name": "Rex",
  "species": "Canine",
  "breed": "Golden Retriever",
  "weight": 35,
  "age": 4,
  "description": "Animal muito ativo",
  "cpf_tutor": "12345678901"
}
```

#### Possíveis Erros:
- 400 Bad Request: Corpo da requisição ou UUID inválido.
- 500 Internal Server Error: Falha ao atualizar o animal.

---

### 3. Deletar Animal
- **Rota:** `DELETE /animals/:id`
- **Descrição:** Remove um animal do sistema.

#### Parâmetros de URL:
- `id`: UUID do animal.

#### Resposta de Sucesso:
- **Código:** 200 OK
- **Corpo:**
```json
{
  "message": "Animal deletado com sucesso"
}
```

#### Possíveis Erros:
- 400 Bad Request: UUID inválido.
- 500 Internal Server Error: Falha ao deletar o animal.

---

### 4. Adicionar Dosagem
- **Rota:** `POST /dosages`
- **Descrição:** Adiciona uma nova dosagem para um animal.

#### Corpo da Requisição:
```json
{
  "animal_id": "UUID do animal",
  "medication_id": "UUID do medicamento",
  "start_date": "2024-09-01",
  "end_date": "2024-09-10",
  "quantity": 2,
  "dosage": "10ml a cada 12 horas",
  "consultation_id": "UUID da consulta",
  "hospitalization_id": "UUID da hospitalização"
}
```

#### Resposta de Sucesso:
- **Código:** 201 Created
- **Corpo:**
```json
{
  "id": "UUID gerado",
  "animal_id": "UUID do animal",
  "medication_id": "UUID do medicamento",
  "start_date": "2024-09-01",
  "end_date": "2024-09-10",
  "quantity": 2,
  "dosage": "10ml a cada 12 horas",
  "consultation_id": "UUID da consulta",
  "hospitalization_id": "UUID da hospitalização"
}
```

#### Possíveis Erros:
- 400 Bad Request: Corpo da requisição inválido.
- 500 Internal Server Error: Falha ao salvar a dosagem.

---

### 5. Adicionar Consulta
- **Rota:** `POST /consultations`
- **Descrição:** Adiciona uma nova consulta para um animal.

#### Corpo da Requisição:
```json
{
  "animal_id": "UUID do animal",
  "crvm": "CRVM do veterinário",
  "consultation_date": "2024-09-01",
  "reason": "Verificação de saúde",
  "observation": "Nenhuma observação"
}
```

#### Resposta de Sucesso:
- **Código:** 201 Created
- **Corpo:**
```json
{
  "message": "Consulta adicionada com sucesso"
}
```

#### Possíveis Erros:
- 400 Bad Request: Corpo da requisição ou formato de data inválido.
- 500 Internal Server Error: Falha ao salvar a consulta.

---

### 6. Adicionar Medicamento
- **Rota:** `POST /medications`
- **Descrição:** Adiciona um novo medicamento no sistema.

#### Corpo da Requisição:
```json
{
  "name": "Ibuprofeno",
  "active_principles": ["Ibuprofeno"],
  "manufacturer": "MedPharma",
  "concentration": "200mg",
  "presentation": "Comprimido",
  "quantity": 100,
  "expiration_date": "2025-12-01"
}
```

#### Resposta de Sucesso:
- **Código:** 201 Created
- **Corpo:**
```json
{
  "id": "UUID gerado",
  "name": "Ibuprofeno",
  "active_principles": ["Ibuprofeno"],
  "manufacturer": "MedPharma",
  "concentration": "200mg",
  "presentation": "Comprimido",
  "quantity": 100,
  "expiration_date": "2025-12-01"
}
```

#### Possíveis Erros:
- 400 Bad Request: Corpo da requisição inválido ou formato de data inválido.
- 500 Internal Server Error: Falha ao salvar o medicamento.

---

### 7. Deletar Medicamento
- **Rota:** `DELETE /medications/:id`
- **Descrição:** Remove um medicamento do sistema.

#### Parâmetros de URL:
- `id`: UUID do medicamento.

#### Resposta de Sucesso:
- **Código:** 200 OK
- **Corpo:**
```json
{
  "message": "Medicamento deletado com sucesso"
}
```

#### Possíveis Erros:
- 400 Bad Request: UUID inválido.
- 500 Internal Server Error: Falha ao deletar o medicamento.