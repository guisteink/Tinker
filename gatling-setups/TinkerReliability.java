package tinker;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import java.time.Duration;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

public class TinkerReliability extends Simulation {

  HttpProtocolBuilder httpProtocol = http
      .baseUrl("http://localhost:3000")
      .acceptHeader("application/json")
      .contentTypeHeader("application/json");

  // Scenario
  ScenarioBuilder scenario = scenario("Health Check")
      .exec(http("Health Check")
          .get("/health")
          .check(status().is(200)));

  // Configure the user load
  {
    setUp(
        scenario.injectOpen(
            rampUsersPerSec(1).to(500).during(Duration.ofMinutes(10)), // Aumenta gradualmente de 1 para 500 usuários
                                                                       // por segundo durante 10 minutos
            constantUsersPerSec(500).during(Duration.ofMinutes(30)) // Mantém 500 usuários por segundo por 30 minutos
        )).protocols(httpProtocol);
  }
}
